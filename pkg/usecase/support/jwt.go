package support

import (
	"crypto/rand"
	"encoding/base64"
	"os"
	e "sample-gin-ddd/pkg/errors"
	"sample-gin-ddd/pkg/model"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func CreateState() string {
	b := make([]byte, 32)
	if _, err := rand.Read(b); err != nil {
		return ""
	}
	return base64.URLEncoding.EncodeToString(b)
}

func CreateJwt(account *model.Accounts, state string) (string, error) {
	type claim struct {
		AccountID string `json:"account_id" binding:"required"`
		Authority string `json:"authority" binding:"required"`
		State     string `json:"state" binding:"required"`
		jwt.RegisteredClaims
	}

	claims := &claim{
		account.AccountID,
		account.Authority,
		state,
		jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 2)),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	t, err := token.SignedString([]byte(os.Getenv("JWT_SECRET_KEY")))
	if err != nil {
		return "", e.INTERNAL_SERVER_ERROR
	}

	return t, nil
}
