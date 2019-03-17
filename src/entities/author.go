package entities

import "gopkg.in/mgo.v2/bson"

// Author struct
type Author struct {
	ID        bson.ObjectId `bson:"_id"`
	Firstname string        `bson:"firstname"`
	Lastname  string        `bson:"lastname"`
}
