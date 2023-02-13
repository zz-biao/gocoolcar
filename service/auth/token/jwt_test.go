package token

import (
	"github.com/dgrijalva/jwt-go"
	"testing"
	"time"
)

const privateKey = ``

func TestGenerateToken(t *testing.T) {
	key, err := jwt.ParseRSAPrivateKeyFromPEM([]byte(privateKey))
	if err != nil {
		t.Fatalf("cannot parse private key: %v", err)
	}
	g := NewJWTTokenGen("coolcar/auth", key)
	g.nowFunc = func() time.Time {
		return time.Unix(1516239022, 0)
	}
	token, err := g.GenerateToken("123", 2*time.Hour)
	if err != nil {
		t.Errorf("cannot generate token: %v", err)
	}
	want := ""
	if token != want {
		t.Errorf("wrong token generate. want: %q,got %q", want, token)
	}
}
