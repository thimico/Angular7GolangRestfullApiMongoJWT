package entities

import "gopkg.in/mgo.v2/bson"

// Author struct
type Author struct {
	Id        bson.ObjectId `bson:"_id"`
	Firstname string        `bson:"firstname"`
	Lastname  string        `bson:"lastname"`
}
