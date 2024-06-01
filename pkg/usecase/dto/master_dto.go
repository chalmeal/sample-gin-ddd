package dto

import "sample-gin-ddd/pkg/model"

type GetAccountDto struct {
	Account *model.Accounts
}

type EditAccountDto struct {
	Message string
}

type DeleteAccountDto struct {
	Message string
}
