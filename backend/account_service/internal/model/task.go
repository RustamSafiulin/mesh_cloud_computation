package model

import (
	"gopkg.in/mgo.v2/bson"
)

const (
	StateCreated = 0
	StateRunning = 1
	StateStopped = 2
)

type Task struct {
	ID          bson.ObjectId `bson:"_id,omitempty"`
	AccountID   bson.ObjectId `bson:"account_id,omitempty"`
	Description string		  `bson:"description,omitempty"`
	CreatedAt   int64         `bson:"created_at,omitempty"`
	StartedAt   int64         `bson:"started_at,omitempty"`
	CompletedAt int64 		  `bson:"completed_at,omitempty"`
	State       int			  `bson:"state,omitempty"`
}

