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

	//表格验证
	cases := []struct {
		name    string
		tkn     string
		now     time.Time
		want    string
		wantErr bool
	}{
		{
			name: "valid_token",
			tkn:  "1111",
			now:  time.Unix(1516239022, 0),
			want: "11111",
		},
		{
			name:    "valid_token1",
			tkn:     "1111",
			now:     time.Unix(1517239022, 0),
			wantErr: true,
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			jwt.TimeFunc = func() time.Time {
				return c.now
			}
			accountID, err := v.Verifier(c.tkn)

			if !c.wantErr && err != nil {
				t.Errorf("verification faild: %v", err)
			}
			if c.wantErr && err == nil {
				t.Errorf("want error got no error")
			}
			if accountID != c.want {
				t.Errorf("want account id want: %q ,got :%q", c.want, accountID)
			}
		})
	}

}
