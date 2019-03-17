package models

import (
	"../entities"
	"golang.org/x/crypto/bcrypt"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type AccountModel struct {
	DB *mgo.Database
}

func (accountModel AccountModel) Create(account *entities.Account) error {
	hash, _ := bcrypt.GenerateFromPassword([]byte(account.Password), bcrypt.DefaultCost)
	account.Password = string(hash)
	return accountModel.DB.C("account").Insert(&account)
}

func (accountModel AccountModel) CheckUsernameAndPassword(username, password string) bool {
	var account entities.Account
	err := accountModel.DB.C("account").Find(bson.M{
		"username": username,
	}).One(&account)
	if err != nil {
		return false
	} else {
		return bcrypt.CompareHashAndPassword([]byte(account.Password), []byte(password)) == nil
	}
}
