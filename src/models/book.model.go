package models

import (
	"../config"
	"../entities"
)

type ProductModel struct {

}

func (productModel ProductModel) FindAll() ([]entities.Product, error) {
	db, err := config.GetDB()
	if err != nil {
		return nil, err
	} else {
		var products []entities.Product
		db.Find(&products)
		return products, nil
	}
}

func (productModel ProductModel) Search(keyword string) ([]entities.Product, error) {
	db, err := config.GetDB()
	var products []entities.Product
	db.Where("name like ?", "%"+keyword+"%").Find(&products)
	return products, err
}

func (productModel ProductModel) Find(id string) (entities.Product, error) {
	db, err := config.GetDB()
	var product entities.Product
	db.Where("id = ?", id).Find(&product)
	return product, err
}

func (productModel ProductModel) Create(product *entities.Product) error {
	db, err := config.GetDB()
	db.Create(&product)
	return err
}

func (productModel ProductModel) Delete(product entities.Product) error {
	db, err := config.GetDB()
	db.Delete(product)
	return err
}

func (productModel ProductModel) Update(id string, product *entities.Product) error {
	db, err := config.GetDB()
	db.Table("product").Where("id = ?", id).Updates(&product)
	return err
}

func (productModel ProductModel) Save(product *entities.Product) error {
	db, err := config.GetDB()
	db.Save(&product)
	return err
}