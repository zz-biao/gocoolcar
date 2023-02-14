package token

import (
	"github.com/dgrijalva/jwt-go"
	"testing"
	"time"
)

const pubilcKey = ""

func TestVerifier(t *testing.T) {
	pubKey, err := jwt.ParseRSAPublicKeyFromPEM([]byte(pubilcKey))
	if err != nil {
		t.Errorf("cannot parse public key: %v", err)
	}
	v := &JWTTokenVerifier{
		PublicKey: pubKey,
	}

	tkn := "ssss"

	//固定jwt时间
	jwt.TimeFunc = func() time.Time {
		return time.Unix(1516239022, 0)
	}

	accountID, err := v.Verifier(tkn)
	want := "11111"
	if err != nil {
		t.Errorf("wrong account id. want: %q, got: %q", want, accountID)
	}
}
