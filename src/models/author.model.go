package models

import (
	"../config"
	"../entities"
)

type BookModel struct {

}

func (bookModel BookModel) FindAll() ([]entities.Book, error) {
	db, err := config.GetDB()
	if err != nil {
		return nil, err
	} else {
		var books []entities.Book
		db.Find(&books)
		return books, nil
	}
}

func (bookModel BookModel) Search(keyword string) ([]entities.Book, error) {
	db, err := config.GetDB()
	var books []entities.Book
	db.Where("title like ?", "%"+keyword+"%").Find(&books)
	return books, err
}

func (bookModel BookModel) Find(id string) (entities.Book, error) {
	db, err := config.GetDB()
	var book entities.Book
	db.Where("id = ?", id).Find(&book)
	return book, err
}

func (bookModel BookModel) Create(book *entities.Book) error {
	db, err := config.GetDB()
	db.Create(&book)
	return err
}

func (bookModel BookModel) Delete(book entities.Book) error {
	db, err := config.GetDB()
	db.Delete(book)
	return err
}

func (bookModel BookModel) Update(id string, book *entities.Book) error {
	db, err := config.GetDB()
	db.Table("book").Where("id = ?", id).Updates(&book)
	return err
}

func (bookModel BookModel) Save(book *entities.Book) error {
	db, err := config.GetDB()
	db.Save(&book)
	return err
}