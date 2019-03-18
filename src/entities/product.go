package entities

import "gopkg.in/mgo.v2/bson"

// Product struct (Model)
type Product struct {
	Id     bson.ObjectId `json:"id" bson:"id"`
	Name   string        `json:"name" bson:"name"`
	Price  float64       `json:"price" bson:"price"`
	Status bool          `json:"status" bson:"status"`
}

func (e *Product) New() IEntity {
	return e
}
