package dto

import (
	"github.com/RustamSafiulin/mesh_cloud_computation/backend/account_service/internal/model"
)

type AccountDto struct {
	AccountID string `json:"id,omitempty"`
	Username  string `json:"username,omitempty"`
	Email     string `json:"email,omitempty"`
}

func AccountDtoFromAccount(account *model.Account) *AccountDto {

	accountDto := &AccountDto{
		AccountID: account.ID.Hex(),
		Username:  account.Name,
		Email:     account.Email,
	}

	return accountDto
}

func AccountDtoListFromAccountList(accounts []model.Account) []AccountDto {

	var result []AccountDto
	for _, account := range accounts {
		accountDto := AccountDtoFromAccount(&account)
		result = append(result, *accountDto)
	}

	return result
}
