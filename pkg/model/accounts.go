package model

import (
	e "sample-gin-ddd/pkg/errors"

	"sample-gin-ddd/pkg/util"
	"time"

	"sample-gin-ddd/pkg/infrastracture/security"
)

var (
	sec = security.NewSecurity()
)

type Accounts struct {
	AccountID string     `gorm:"primaryKey; type:varchar(50); not null;unique" json:"account_id"`
	Password  string     `gorm:"type:varchar(500)" json:"password"`
	Name      string     `gorm:"type:varchar(50)" json:"name"`
	Email     string     `gorm:"type:varchar(50)" json:"email"`
	AvatorUrl string     `gorm:"type:varchar(200)" json:"avator_url"`
	Authority string     `gorm:"type:varchar(20); not null" json:"authority"`
	AuthType  string     `gorm:"type:varchar(20); not null" json:"auth_type"`
	CreatedAt time.Time  `gorm:"type:datetime; not null" json:"created_at"`
	UpdatedAt *time.Time `gorm:"type:datetime; autoUpdateTime:false" json:"updated_at"`
	DeletedAt *time.Time `gorm:"type:datetime" json:"deleted_at"`
}

func NewAccount() *Accounts {
	return &Accounts{}
}

type VerLogin struct {
	AccountID string `json:"account_id" binding:"required"`
	Password  string `json:"password" binding:"required"`
}

func (account *Accounts) VerifyLogin(a *Accounts, p *VerLogin) error {
	if a.AccountID == "" || a.Password != sec.Hash(p.Password) {
		return e.LOGIN_UN_AUTHORIZATION
	}
	return nil
}

func (account *Accounts) HasEmail(ac *Accounts) error {
	if ac.Email != "" {
		return e.TEMPORARY_REGISTER_ACCOUNT_ALREADY
	}
	return nil
}

func (account *Accounts) IsDelete(ac *Accounts) error {
	if !ac.DeletedAt.IsZero() {
		return e.DELETE_ACCOUNT_ALREADY
	}
	return nil
}

type SpecifyAccount struct {
	AccountID string `json:"account_id"`
}

func (account *Accounts) Specify(accountId string) *Accounts {
	return &Accounts{
		AccountID: accountId,
	}
}

type FindAccount struct {
	AccountID string `json:"account_id"`
	Name      string `json:"name"`
	Authority string `json:"authority"`
	AuthType  string `json:"auth_type"`
}

func (account *Accounts) Find(p *FindAccount) *Accounts {
	return &Accounts{
		AccountID: p.AccountID,
		Name:      p.Name,
		Authority: p.Authority,
		AuthType:  p.AuthType,
	}
}

type TmpAccount struct {
	Email string `json:"email" binding:"required"`
}

func (account *Accounts) Temporary(p *TmpAccount) *Accounts {
	return &Accounts{
		AccountID: util.EmailToId(p.Email),
		Email:     p.Email,
		Authority: "TEMPORARY",
		AuthType:  "APP",
	}
}

type RegAccount struct {
	AccountID string `json:"account_id" binding:"required"`
	Password  string `json:"password" binding:"required"`
	Name      string `json:"name" binding:"required"`
	AvatorUrl string `json:"avator_url"`
}

func (account *Accounts) Regist(p *RegAccount) *Accounts {
	return &Accounts{
		Password:  sec.Hash(p.Password),
		Name:      p.Name,
		AvatorUrl: p.AvatorUrl,
		Authority: "NORMAL",
	}
}

type EdtAccount struct {
	Name      string `json:"name"`
	Email     string `json:"email"`
	AvatorUrl string `json:"avator_url"`
}

func (account *Accounts) Edit(p *EdtAccount) *Accounts {
	return &Accounts{
		Name:      p.Name,
		Email:     p.Email,
		AvatorUrl: p.AvatorUrl,
		UpdatedAt: util.NowDateTime(),
	}
}
