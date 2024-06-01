package db

import (
	"sample-gin-ddd/pkg/usecase/dto"

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
		} else if tx.Commit().Error != nil {
			data.Error = tx.Commit().Error
		} else {
			tx.Commit()
			return
		}
	}()

	return txFunc(tx)
}
