package dao

import (
	"../entities"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type BookDAO struct {
	DB *mgo.Database
}

func (bookDAO BookDAO) FindAll() ([]entities.Book, error) {
	var books []entities.Book
	err := bookDAO.DB.C("book").Find(bson.M{}).All(&books)
	if err != nil {
		return nil, err
	} else {
		return books, nil
	}
}

func (bookDAO BookDAO) Find(id string) (entities.Book, error) {
	var book entities.Book
	err := bookDAO.DB.C("book").FindId(bson.ObjectIdHex(id)).One(&book)
	return book, err
}

func (bookDAO BookDAO) Create(book *entities.Book) error {
	return bookDAO.DB.C("book").Insert(&book)
}

func (bookDAO BookDAO) Delete(id string) error {
	return bookDAO.DB.C("book").RemoveId(bson.ObjectIdHex(id))
}

func (bookDAO BookDAO) Update(book *entities.Book) error {
	return bookDAO.DB.C("book").UpdateId(book.Id, book)
}
