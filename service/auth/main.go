package main

import (
	"context"
	authpb "coolcar/auth/api/gen/v1"
	"coolcar/auth/auth"
	"coolcar/auth/dao"
	"coolcar/auth/token"
	"coolcar/auth/wechat"
	"github.com/dgrijalva/jwt-go"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"io/ioutil"
	"log"
	"net"
	"os"
	"time"
)

func main() {
	logger, err := zap.NewDevelopment()
	if err != nil {
		log.Fatalf("cannot create logger: %v", err)
	}
	lis, err := net.Listen("tcp", ":8081")
	if err != nil {
		logger.Fatal("cannot listen", zap.Error(err))
	}

	c := context.Background()
	mongoCilent, err := mongo.Connect(c, options.Client().ApplyURI("mongodb://root:123456@localhost:27017/?authSource=admin"))
	if err != nil {
		logger.Fatal("cannot connect mongodb", zap.Error(err))
	}

	open, err := os.Open("auth/private.key")
	if err != nil {
		logger.Fatal("cannot open private key ", zap.Error(err))
	}
	pkBytes, err := ioutil.ReadAll(open)
	if err != nil {
		logger.Fatal("cannot read private key ", zap.Error(err))
	}
	key, err := jwt.ParseRSAPrivateKeyFromPEM(pkBytes)
	if err != nil {
		return
	}

	s := grpc.NewServer()
	authpb.RegisterAuthServiceServer(s, &auth.Service{ //设置配置
		OpenIDResolver: &wechat.Service{
			AppID:     "sss",
			AppSecret: "222",
		},
		Mongo:          dao.NewMongo(mongoCilent.Database("coolcar")),
		Logger:         logger,
		TokenExpire:    2 * time.Hour,
		TokenGenerator: token.NewJWTTokenGen("coolcar/auth", key),
	})
	err = s.Serve(lis)
	if err != nil {
		logger.Fatal("cannot server", zap.Error(err))
	}
}
