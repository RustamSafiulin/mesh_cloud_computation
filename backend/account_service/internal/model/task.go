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
	DataUrl     string 		  `bson:"data_url,omitempty"`
	CreatedAt   int64        `bson:"created_at,omitempty"`
	StartedAt   int64        `bson:"started_at,omitempty"`
	CompletedAt int64 		  `bson:"completed_at,omitempty"`
	State       int			  `bson:"state,omitempty"`
}

type TaskDto struct {
	ID 		    string `json:"task_id,omitempty"`
	AccountID   string `json:"account_id,omitempty"`
	Description string `json:"description,omitempty"`
	DataUrl	    string `json:"data_url,omitempty"`
	StartedAt   int64 `json:"started_at,omitempty"`
	CompletedAt int64 `json:"completed_at,omitempty"`
	State       int    `json:"state,omitempty"`
}

