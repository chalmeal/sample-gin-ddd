package repository

import (
	"log"
	e "sample-gin-ddd/pkg/errors"
	"sample-gin-ddd/pkg/model"
	"sample-gin-ddd/pkg/util"

	"gorm.io/gorm"
)

type AccountRepository interface {
	Get(tx *gorm.DB, query interface{}) (*model.Accounts, error)
	Find(tx *gorm.DB) (*[]model.Accounts, error)
	Save(tx *gorm.DB, param *model.Accounts) error
	Update(tx *gorm.DB, query interface{}, param *model.Accounts) error
	Delete(tx *gorm.DB, query interface{}) error
}

type AccountRepositoryImpl struct{}

func NewAccountRepository() AccountRepository {
	return &AccountRepositoryImpl{}
}

func (rep *AccountRepositoryImpl) Get(tx *gorm.DB, query interface{}) (*model.Accounts, error) {
	var account model.Accounts
	err := tx.Where(query).First(&account)
	if err.Error != nil {
		return &model.Accounts{}, e.GET_ACCOUNT_NOT_FOUND
	}
	return &account, nil
}

func (rep *AccountRepositoryImpl) Find(query *gorm.DB) (*[]model.Accounts, error) {
	var accounts []model.Accounts
	err := query.Find(&accounts)
	if err.Error != nil {
		log.Println(err.Error)
		return &[]model.Accounts{}, e.INTERNAL_SERVER_ERROR
	}
	return &accounts, nil
}

func (rep *AccountRepositoryImpl) Save(tx *gorm.DB, param *model.Accounts) error {
	if err := tx.Create(param).Error; err != nil {
		log.Println(err)
		return e.INTERNAL_SERVER_ERROR
	}
	return nil
}

func (rep *AccountRepositoryImpl) Update(tx *gorm.DB, query interface{}, param *model.Accounts) error {
	var account model.Accounts
	record := tx.Where(query).First(&account)
	if record.Error != nil || account.AccountID == "" {
		log.Println(record.Error)
		return e.GET_ACCOUNT_NOT_FOUND
	}
	if err := record.Updates(param).Error; err != nil {
		log.Println(err)
		return e.INTERNAL_SERVER_ERROR
	}
	return nil
}

func (rep *AccountRepositoryImpl) Delete(tx *gorm.DB, query interface{}) error {
	var account model.Accounts
	record := tx.Where(query).First(&account)
	if record.Error != nil || util.IsEmpty(account, model.Accounts{}) {
		log.Println(record.Error)
		return e.GET_ACCOUNT_NOT_FOUND
	}
	if err := record.Delete(&account).Error; err != nil {
		log.Println(err)
		return e.INTERNAL_SERVER_ERROR
	}
	return nil
}
