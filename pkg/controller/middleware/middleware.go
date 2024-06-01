package middleware

import (
	"fmt"
	"net/http"
	"os"

	"sample-gin-ddd/pkg/controller/authority"
	e "sample-gin-ddd/pkg/errors"
	"sample-gin-ddd/pkg/infrastracture/session"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

type Middleware struct {
	sess      session.SessionInfoImpl
	authority authority.Authority
}

func NewMiddleware() *Middleware {
	return &Middleware{
		sess:      session.NewSessionInfo(),
		authority: *authority.NewAuthority(),
	}
}

func (mw *Middleware) AuthMiddleware(permit []string) func(c *gin.Context) {
	return func(c *gin.Context) {
		// アクセストークンの検証
		{
			err := mw.validateAccessToken(c)
			if err != nil {
				mw.sess.Clear(c)
				c.JSON(http.StatusUnauthorized, gin.H{
					"Error": err.Error(),
				})
				c.Abort()
				return
			}
		}
		// APIアクセス権限の検証
		{
			err := mw.authority.PreAuthorize(c, permit)
			if err != nil {
				mw.sess.Clear(c)
				c.JSON(http.StatusUnauthorized, gin.H{
					"Error": err.Error(),
				})
				c.Abort()
				return
			}
		}
	}
}

func (mw *Middleware) validateAccessToken(c *gin.Context) error {
	token := c.GetHeader("Authorization")
	t, err := jwt.Parse(token, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			mw.sess.Clear(c)
			return nil, fmt.Errorf("unexpected signing method: %s", t.Header["alg"])
		}
		return []byte(os.Getenv("JWT_SECRET_KEY")), nil
	})
	if err != nil {
		return e.ACCESS_TOKEN_VERIFY_FAILURE
	}

	if !t.Valid {
		return e.ACCESS_TOKEN_VERIFY_FAILURE
	}

	claims, ok := t.Claims.(jwt.MapClaims)
	if !ok {
		return e.ACCESS_TOKEN_VERIFY_FAILURE
	}

	state := mw.sess.GetState(c)
	if state != claims["state"] {
		return e.ACCESS_TOKEN_VERIFY_FAILURE
	}

	return nil
}
