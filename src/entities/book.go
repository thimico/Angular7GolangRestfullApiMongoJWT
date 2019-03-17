package entities

import "gopkg.in/mgo.v2/bson"

// Book struct (Model)
type Book struct {
	ID     bson.ObjectId `bson:"_id"`
	Isbn   string        `bson:"isbn"`
	Title  string        `bson:"title"`
	Price  float64       `bson:"price"`
	Status bool          `bson:"status"`
	Author *Author       `bson:"author"`
}
