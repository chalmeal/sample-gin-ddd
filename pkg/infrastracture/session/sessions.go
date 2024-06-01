package session

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

type SessionInfoImpl interface {
	GetAccountId(c *gin.Context) string
	GetAuthority(c *gin.Context) string
	GetState(c *gin.Context) string
	Set(c *gin.Context, s *SessionInfo)
	Clear(c *gin.Context)
}

type SessionInfo struct {
	AccountId string
	Authority string
	State     string
}

func NewSessionInfo() *SessionInfo {
	return &SessionInfo{}
}

func (si *SessionInfo) GetAccountId(c *gin.Context) string {
	sess := sessions.Default(c)
	if sess.Get("account_id") == nil {
		return ""
	}
	return sess.Get("account_id").(string)
}

func (si *SessionInfo) GetAuthority(c *gin.Context) string {
	sess := sessions.Default(c)
	if sess.Get("authority") == nil {
		return ""
	}
	return sess.Get("authority").(string)
}

func (si *SessionInfo) GetState(c *gin.Context) string {
	sess := sessions.Default(c)
	if sess.Get("state") == nil {
		return ""
	}
	return sess.Get("state").(string)
}

func (si *SessionInfo) Set(c *gin.Context, s *SessionInfo) {
	sess := sessions.Default(c)
	sess.Set("account_id", s.AccountId)
	sess.Set("authority", s.Authority)
	sess.Set("state", s.State)
	sess.Save()
}

func (si *SessionInfo) Clear(c *gin.Context) {
	sess := sessions.Default(c)
	sess.Clear()
	sess.Save()
}
