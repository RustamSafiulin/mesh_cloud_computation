package storage

import (
	"github.com/RustamSafiulin/mesh_cloud_computation/backend/account_service/cmd"
	"github.com/RustamSafiulin/mesh_cloud_computation/backend/account_service/internal/model"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

var tasksCollectionName = "tasks"

type TaskStorage struct {
	mgoSession *mgo.Session
}

func NewTaskStorage(mgoSession *mgo.Session) *TaskStorage {
	return &TaskStorage{mgoSession: mgoSession}
}

func (storage *TaskStorage) collection() *mgo.Collection {
	return storage.mgoSession.DB(cmd.DbName).C(tasksCollectionName)
}

func (storage *TaskStorage) FindById(id string) (*model.Task, error) {
	var task *model.Task
	err := storage.collection().FindId(bson.ObjectIdHex(id)).One(&task)
	return task, err
}

func (storage *TaskStorage) Insert(t *model.Task) error {
	t.ID = bson.NewObjectId()
	return storage.collection().Insert(t)
}

func (storage *TaskStorage) FindAll(accountId string) ([]model.Task, error) {

	query := bson.M{
		"account_id": bson.M{
			"$eq": bson.ObjectIdHex(accountId),
		},
	}

	var tasks []model.Task
	err := storage.collection().Find(query).All(&tasks)
	return tasks, err
}

func (storage *TaskStorage) Delete(id string) error {
	return storage.collection().RemoveId(bson.ObjectIdHex(id))
}

func (storage *TaskStorage) Update(t *model.Task) error {
	return storage.collection().UpdateId(t.ID, t)
}