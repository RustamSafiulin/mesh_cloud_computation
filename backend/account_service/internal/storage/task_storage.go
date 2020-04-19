package storage

import (
	"github.com/RustamSafiulin/mesh_cloud_computation/backend/account_service/cmd"
	"github.com/RustamSafiulin/mesh_cloud_computation/backend/account_service/internal/model"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

var tasksCollectionName = "tasks"
var taskFilesCollectionName = "task_files"

type TaskStorage struct {
	mgoSession *mgo.Session
}

func NewTaskStorage(mgoSession *mgo.Session) *TaskStorage {
	return &TaskStorage{mgoSession: mgoSession}
}

func (storage *TaskStorage) tasksCollection() *mgo.Collection {
	return storage.mgoSession.DB(cmd.DbName).C(tasksCollectionName)
}

func (storage *TaskStorage) taskFilesCollection() *mgo.Collection {
	return storage.mgoSession.DB(cmd.DbName).C(taskFilesCollectionName)
}

func (storage *TaskStorage) FindById(id string) (*model.Task, error) {
	var task *model.Task
	err := storage.tasksCollection().FindId(bson.ObjectIdHex(id)).One(&task)
	return task, err
}

func (storage *TaskStorage) Insert(t *model.Task) (*model.Task, error) {
	t.ID = bson.NewObjectId()
	err := storage.tasksCollection().Insert(t)
	return t, err
}

func (storage *TaskStorage) FindAllByAccount(accountId string) ([]model.Task, error) {

	query := bson.M{
		"account_id": bson.M{
			"$eq": bson.ObjectIdHex(accountId),
		},
	}

	var tasks []model.Task
	err := storage.tasksCollection().Find(query).All(&tasks)
	return tasks, err
}

func (storage *TaskStorage) Delete(id string) error {
	return storage.tasksCollection().RemoveId(bson.ObjectIdHex(id))
}

func (storage *TaskStorage) Update(t *model.Task) error {
	return storage.tasksCollection().UpdateId(t.ID, t)
}

func (storage *TaskStorage) InsertTaskFile(tf *model.TaskFile) (*model.TaskFile, error) {
	tf.ID = bson.NewObjectId()
	err := storage.taskFilesCollection().Insert(tf)
	return tf, err
}

func (storage *TaskStorage) DeleteTaskFile(taskFileId string) error {
	return storage.taskFilesCollection().RemoveId(bson.ObjectIdHex(taskFileId))
}