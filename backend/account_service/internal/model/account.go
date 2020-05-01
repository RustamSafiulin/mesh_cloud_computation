package model

import (
	"gopkg.in/mgo.v2/bson"
)

type Account struct {
	ID           bson.ObjectId `bson:"_id,omitempty"`
	Name     	 string 	   `bson:"name,omitempty"`
	Email    	 string 	   `bson:"email,omitempty"`
	PasswordHash string        `bson:"password_hash,omitempty"`
	CreatedAt    int64         `bson:"created_at"`
}



