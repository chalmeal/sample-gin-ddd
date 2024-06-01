package usecase

import (
	e "sample-gin-ddd/pkg/errors"
	"sample-gin-ddd/pkg/infrastracture/db"
	"sample-gin-ddd/pkg/infrastracture/repository"
	"sample-gin-ddd/pkg/infrastracture/session"
	"sample-gin-ddd/pkg/model"
	"sample-gin-ddd/pkg/usecase/dto"
	"sample-gin-ddd/pkg/usecase/support"

	"gorm.io/gorm"
)

type AppService struct {
	db      *gorm.DB
	sess    session.SessionInfoImpl
	rep     repository.AccountRepository
	account model.Accounts
}

func NewAppService() *AppService {
	return &AppService{
		db:      db.GetDB(),
		sess:    session.NewSessionInfo(),
		rep:     repository.NewAccountRepository(),
		account: *model.NewAccount(),
	}
}

// アカウントにログインします。
func (app *AppService) Login(param *model.VerLogin) (*dto.Dto, *session.SessionInfo) {
	var sess *session.SessionInfo
	dto := db.Tx(app.db, func(tx *gorm.DB) *dto.Dto {
		query := &model.Accounts{
			AccountID: param.AccountID,
		}
		a, _ := app.rep.Get(tx, query)
		if err := app.account.VerifyLogin(a, param); err != nil {
			return &dto.Dto{
				Error: err,
			}
		}

		state := support.CreateState()
		token, err := support.CreateJwt(a, state)
		if err != nil {
			return &dto.Dto{
				Error: err,
			}
		}

		sess = &session.SessionInfo{
			AccountId: a.AccountID,
			Authority: a.Authority,
			State:     state,
		}

		return &dto.Dto{
			Result: &dto.LoginDto{
				AccessToken: token,
			},
			Error: nil,
		}
	})

	return dto, sess
}

// アカウントを仮登録します。
func (master *AppService) RegisterTemporaryAccount(param *model.TmpAccount) *dto.Dto {
	return db.Tx(master.db, func(tx *gorm.DB) *dto.Dto {
		query := &model.Accounts{
			Email: param.Email,
		}
		ac, _ := master.rep.Get(tx, query)
		if err := master.account.HasEmail(ac); err != nil {
			return &dto.Dto{
				Error: err,
			}
		}

		tmp := master.account.Temporary(param)
		if err := master.rep.Save(tx, tmp); err != nil {
			return &dto.Dto{
				Error: err,
			}
		}

		if err := support.SendRegisterMail(param.Email); err != nil {
			return &dto.Dto{
				Error: err,
			}
		}

		return &dto.Dto{
			Result: &dto.RegisterTemporaryAccountDto{
				Message: e.TEMPORARY_ACCOUNT_SUCCESS.Error(),
			},
			Error: nil,
		}
	})
}

// アカウントを新規登録します。
func (master *AppService) RegisterAccount(param *model.RegAccount) *dto.Dto {
	return db.Tx(master.db, func(tx *gorm.DB) *dto.Dto {
		query := &model.Accounts{
			AccountID: param.AccountID,
		}
		reg := master.account.Regist(param)
		if err := master.rep.Update(tx, query, reg); err != nil {
			return &dto.Dto{
				Error: err,
			}
		}

		return &dto.Dto{
			Result: &dto.RegisterAccountDto{
				Message: e.REGISTER_ACCOUNT_SUCCESS.Error(),
			},
			Error: nil,
		}
	})
}
