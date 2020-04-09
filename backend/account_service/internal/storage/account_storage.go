package storage

import (
	"github.com/RustamSafiulin/mesh_cloud_computation/backend/account_service/cmd"
	"github.com/RustamSafiulin/mesh_cloud_computation/backend/account_service/internal/model"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

var accountCollectionName = "accounts"

type AccountStorage struct {
	mgoSession *mgo.Session
}

func NewAccountStorage(mgoSession *mgo.Session) *AccountStorage {
	return &AccountStorage{mgoSession: mgoSession}
}

func (storage *AccountStorage) collection() *mgo.Collection {
	return storage.mgoSession.DB(cmd.DbName).C(accountCollectionName)
}

func (storage *AccountStorage) FindById(id string) (*model.Account, error) {

	var account *model.Account

	collection := storage.collection()
	query := bson.M{
		"_id": bson.M{
			"$eq": bson.ObjectIdHex(id),
		},
	}

	err := collection.Find(query).One(account)
	return account, err
}

func (storage *AccountStorage) FindByEmail(email string) (*model.Account, error) {

	var account *model.Account

	collection := storage.collection()
	query := bson.M{
		"email": bson.M{
			"$eq": email,
		},
	}

	err := collection.Find(query).One(&account)
	return account, err
}

func (storage *AccountStorage) Insert(account *model.Account) error {
	return storage.collection().Insert(account)
}

func (storage *AccountStorage)Update(account *model.Account) error {
	return storage.collection().UpdateId(account.ID, account)
}