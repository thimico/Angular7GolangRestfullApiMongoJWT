package models

import (
	"../entities"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type BookModel struct {
	DB *mgo.Database
}

func (bookModel BookModel) Create(book *entities.Book) error {
	return bookModel.DB.C("book").Insert(&book)
}

func (bookModel BookModel) FindAll() ([]entities.Book, error) {
	var books []entities.Book
	err := bookModel.DB.C("book").Find(bson.M{}).All(&books)
	if err != nil {
		return nil, err
	} else {
		return books, nil
	}
}
