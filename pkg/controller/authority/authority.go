package authority

import (
	e "sample-gin-ddd/pkg/errors"
	"sample-gin-ddd/pkg/infrastracture/session"

	"github.com/gin-gonic/gin"
	"golang.org/x/exp/slices"
)

var (
	AUTHORIZE_ADMIN  = []string{"ADMIN"}
	AUTHORIZE_NORMAL = []string{"ADMIN", "NORMAL"}
)

type Authority struct {
	sess session.SessionInfoImpl
}

func NewAuthority() *Authority {
	return &Authority{
		sess: session.NewSessionInfo(),
	}
}

func (authority *Authority) PreAuthorize(c *gin.Context, permit []string) error {
	auth := authority.sess.GetAuthority(c)
	if !slices.Contains(permit, auth) {
		return e.NOT_ACCESS_AUTHORIZE
	}
	return nil
}
