package model

import (
	"gopkg.in/mgo.v2/bson"
)

const (
	StateCreated = 0
	StateRunning = 1
	StateCompleted = 2
	StateCancelled = 3
)

func GetStateStringFromState(state int) string {

	if state == StateCreated {
		return "Created"
	} else if state == StateRunning {
		return "Running"
	} else if state == StateCompleted {
		return "Completed"
	} else if state == StateCancelled {
		return "Cancelled"
	}

	return "Unknown"
}

type Task struct {
	ID            bson.ObjectId `bson:"_id,omitempty"`
	AccountID     bson.ObjectId `bson:"account_id,omitempty"`
	Description   string		`bson:"description,omitempty"`
	CreatedAt     int64         `bson:"created_at"`
	StartedAt     int64         `bson:"started_at"`
	CompletedAt   int64 		`bson:"completed_at"`
	State         int			`bson:"state"`
}

