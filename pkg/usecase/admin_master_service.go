package usecase

import (
	e "sample-gin-ddd/pkg/errors"
	"sample-gin-ddd/pkg/infrastracture/db"
	"sample-gin-ddd/pkg/infrastracture/repository"
	"sample-gin-ddd/pkg/model"
	"sample-gin-ddd/pkg/usecase/dto"

	"gorm.io/gorm"
)

type MasterAdminService struct {
	db      *gorm.DB
	orm     *db.OrmBuilder
	rep     repository.AccountRepository
	account model.Accounts
}

func NewMasterAdminService() *MasterAdminService {
	return &MasterAdminService{
		db:      db.GetDB(),
		orm:     db.NewOrmRepository(),
		rep:     repository.NewAccountRepository(),
		account: *model.NewAccount(),
	}
}

// アカウントを検索します。
func (master *MasterAdminService) FindAccount(param *model.FindAccount) *dto.Dto {
	return db.Tx(master.db, func(tx *gorm.DB) *dto.Dto {
		query := &model.Accounts{
			AccountID: param.AccountID,
			Name:      param.Name,
			Authority: param.Authority,
			AuthType:  param.AuthType,
		}

		// TODO: 完全一致しか取得できていない
		ac, err := master.rep.Find(tx, query)
		if err != nil {
			return &dto.Dto{
				Error: e.INTERNAL_SERVER_ERROR,
			}
		}

		return &dto.Dto{
			Result: &dto.GetAccountAllDto{
				Accounts: ac,
			},
			Error: nil,
		}
	})
}
