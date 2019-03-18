package entities

import (
	"gopkg.in/mgo.v2/bson"
)

type Account struct {
	Id       bson.ObjectId `json:"id" bson:"_id"`
	Username string        `json:"username" bson:"username"`
	Email    string        `json:"email" bson:"email"`
	Password string        `json:"password" bson:"password"`
	FullName string        `json:"fullname" bson:"fullname"`
	Role     string        `json:"role" bson:"role"`
}

func (e *Account) New() IEntity {
	return e
}
