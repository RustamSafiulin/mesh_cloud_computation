package storage

import (
	"github.com/RustamSafiulin/mesh_cloud_computation/backend/account_service/internal/model"
)

type BaseTaskStorage interface {
	FindById(id string) (*model.Task, error)
	Insert(t *model.Task) error
	FindAll(accountId string) ([]model.Task, error)
	Delete(id string) error
	Update(t *model.Task) error
}
