package storage

import (
	"github.com/RustamSafiulin/mesh_cloud_computation/backend/account_service/internal/model"
)

type BaseAccountStorage interface {
	FindById(id string) (*model.Account, error)
	FindByEmail(email string) (*model.Account, error)
	Insert(account *model.Account) error
	Update(account *model.Account) error
}
