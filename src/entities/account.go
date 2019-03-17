package entities

import (
	"gopkg.in/mgo.v2/bson"
)

type Account struct {
	Id       bson.ObjectId `json:"id" bson:"_id"`
	Username string        `json:"username" bson:"username"`
	Password string        `json:"password" bson:"password"`
	FullName string        `json:"fullname" bson:"fullname"`
}
