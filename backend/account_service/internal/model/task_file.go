package model

import "gopkg.in/mgo.v2/bson"

type TaskFile struct {
	ID        bson.ObjectId `bson:"_id,omitempty"`
	TaskID    bson.ObjectId `bson:"task_id,omitempty"`
	Path      string        `bson:"path,omitempty"`
	Name      string        `bson:"name,omitempty"`
	Size      int64         `bson:"size"`
	MD5	      string        `bson:"md5_hash,omitempty"`
	CreatedAt int64			`bson:"created_at"`
}
