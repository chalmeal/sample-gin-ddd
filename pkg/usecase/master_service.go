package usecase

import (
	e "sample-gin-ddd/pkg/errors"

	"sample-gin-ddd/pkg/infrastracture/db"
	"sample-gin-ddd/pkg/infrastracture/repository"
	"sample-gin-ddd/pkg/model"
	"sample-gin-ddd/pkg/usecase/dto"

	"gorm.io/gorm"
)

type MasterService struct {
	db      *gorm.DB
	rep     repository.AccountRepository
	account model.Accounts
}

func NewMasterService() *MasterService {
	return &MasterService{
		db:      db.GetDB(),
		rep:     repository.NewAccountRepository(),
		account: *model.NewAccount(),
	}
}

// 自身のアカウントを取得します。
func (master *MasterService) GetAccount(accountId string) *dto.Dto {
	return db.Tx(master.db, func(tx *gorm.DB) *dto.Dto {
		query := &model.Accounts{
			AccountID: accountId,
		}
		ac, err := master.rep.Get(tx, query)
		if err != nil {
			return &dto.Dto{
				Error: err,
			}
		}

		return &dto.Dto{
			Result: &dto.GetAccountDto{
				Account: ac,
			},
			Error: nil,
		}
	})
}

// 自身のアカウントを編集します。
func (master *MasterService) EditAccount(accountId string, param *model.EdtAccount) *dto.Dto {
	return db.Tx(master.db, func(tx *gorm.DB) *dto.Dto {
		query := &model.SpecifyAccount{
			AccountID: accountId,
		}
		edt := master.account.Edit(param)
		if err := master.rep.Update(tx, query, edt); err != nil {
			return &dto.Dto{
				Error: e.EDIT_ACCOUNT_FAILURE,
			}
		}

		return &dto.Dto{
			Result: &dto.RegisterAccountDto{
				Message: e.EDIT_ACCOUNT_SUCCESS.Error(),
			},
			Error: nil,
		}
	})
}

// 自身のアカウントを物理削除します。
func (master *MasterService) DeleteAccount(accountId string) *dto.Dto {
	return db.Tx(master.db, func(tx *gorm.DB) *dto.Dto {
		query := &model.SpecifyAccount{
			AccountID: accountId,
		}

		if err := master.rep.Delete(tx, query); err != nil {
			return &dto.Dto{
				Error: err,
			}
		}

		return &dto.Dto{
			Result: &dto.DeleteAccountDto{
				Message: e.DELETE_ACCOUNT_SUCCESS.Error(),
			},
			Error: nil,
		}
	})
}
