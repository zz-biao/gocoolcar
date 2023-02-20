package auth

import (
	"context"
	"coolcar/shared/auth/token"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
	"io/ioutil"
	"os"
	"strings"
)

const (
	authorizationHeader = "authorization"
	bearerPrefix        = "Bearer "
)

func Interceptor(publicKey string) (grpc.UnaryServerInterceptor, error) {
	open, err := os.Open(publicKey)
	if err != nil {
		return nil, fmt.Errorf("cannot read public key: %v", err)
	}
	b, err := ioutil.ReadAll(open)
	if err != nil {
		return nil, fmt.Errorf("cannot read public key: %v", err)
	}
	pubKey, err := jwt.ParseRSAPublicKeyFromPEM(b)
	if err != nil {
		return nil, fmt.Errorf("cannot parse public key: %v", err)
	}
	i := &interceptor{
		verifier: &token.JWTTokenVerifier{
			PublicKey: pubKey,
		},
	}
	return i.HandleReq, nil
}

type tokenVerifier interface {
	Verifier(token string) (string, error)
}

type interceptor struct {
	//publicKey *rsa.PublicKey //结构的话引用一般加上地址 * ,interface 从不加
	verifier tokenVerifier
}

func (i *interceptor) HandleReq(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {
	tkn, err := tokenFromContext(ctx)
	if err != nil {
		return nil, status.Error(codes.Unauthenticated, "")
	}
	accId, err := i.verifier.Verifier(tkn)
	if err != nil {
		return nil, status.Errorf(codes.Unauthenticated, "token not verifier: %v", err)
	}

	return handler(ContextWithAccountID(ctx, AccountID(accId)), req)
}

func tokenFromContext(c context.Context) (string, error) {
	unauthenticated := status.Error(codes.Unauthenticated, "")
	m, ok := metadata.FromIncomingContext(c)
	if !ok {
		return "", unauthenticated
	}
	tkn := ""
	for _, v := range m[authorizationHeader] {
		if strings.HasPrefix(v, bearerPrefix) { //检查字符串中含有Bearer
			tkn = v[len(bearerPrefix):] //"Bearer " 以后的字符串
		}
	}
	if tkn == "" {
		return "", unauthenticated
	}
	return tkn, nil
}

type accountIDKey struct {
}

type AccountID string

func (a AccountID) String() string {
	return string(a)
}

func ContextWithAccountID(c context.Context, aid AccountID) context.Context {
	return context.WithValue(c, accountIDKey{}, aid)
}

func AccountIDFromContext(c context.Context) (AccountID, error) {
	v := c.Value(accountIDKey{})
	aid, ok := v.(AccountID)
	if !ok {
		return "", status.Error(codes.Unauthenticated, "")
	}
	return aid, nil
}
