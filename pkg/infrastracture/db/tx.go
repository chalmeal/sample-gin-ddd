package db

import (
	"sample-gin-api/pkg/usecase/dto"

	"gorm.io/gorm"
)

func Tx(db *gorm.DB, txFunc func(*gorm.DB) *dto.Dto) (data *dto.Dto) {
	tx := db.Begin()
	if tx.Error != nil {
		data.Error = tx.Error
		return data
	}

	defer func() {
		if p := recover(); p != nil {
			tx.Rollback()
			return
		}
		if err := tx.Commit().Error; err != nil {
			tx.Rollback()
			data.Error = err
			return
		}
	}()

	return txFunc(tx)
}
