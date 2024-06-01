package security

import (
	"crypto/sha256"
	"fmt"
)

type Security interface {
	Hash(pass string) string
}

type SecurityImpl struct{}

func NewSecurity() Security {
	return &SecurityImpl{}
}

func (sec *SecurityImpl) Hash(key string) string {
	pw := sha256.Sum256([]byte(key))
	return fmt.Sprintf("%x", pw)
}
