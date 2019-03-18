package entities

import "gopkg.in/mgo.v2/bson"

type IEntity interface {
	New() IEntity
}

type AbstractEntity struct {
	Id bson.ObjectId `json:"id" bson:"_id"`
}
