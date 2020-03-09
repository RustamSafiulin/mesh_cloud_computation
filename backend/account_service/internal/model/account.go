package model

import (
	"gopkg.in/mgo.v2/bson"
)

type Account struct {
	ID           bson.ObjectId `bson:"_id,omitempty"`
	Name     	 string 	   `bson:"name,omitempty"`
	Email    	 string 	   `bson:"email,omitempty"`
	PasswordHash string        `bson:"password_hash,omitempty"`
	CreatedAt    int64         `bson:"created_at,omitempty"`
}

type SignupInfoDto struct {
	Username string `json:"username,omitempty"`
	Email 	 string `json:"email,omitempty"`
	Password string `json:"password,omitempty"`
}

type SigninInfoDto struct {
	Email 	 string `json:"email,omitempty"`
	Password string `json:"password,omitempty"`
}

type SessionInfoDto struct {
	AccountID	 string `json:"account_id,omitempty"`
	SessionToken string `json:"session_token,omitempty"`
}

type AccountDto struct {
	AccountID string `json:"account_id,omitempty"`
	Username  string `json:"username,omitempty"`
	Email     string `json:"email,omitempty"`
}

type ErrorMsgResponse struct {
	Error string `json:"error"`
}

type SuccessMsgResponse struct {
	Msg string `json: "msg"`
}



