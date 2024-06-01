package dto

import "sample-gin-ddd/pkg/model"

type GetAccountAllDto struct {
	Accounts *[]model.Accounts
}
