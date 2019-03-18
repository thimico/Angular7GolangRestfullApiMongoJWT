package dao

import (
	"../entities"
	. "./abstractdao"
	"golang.org/x/crypto/bcrypt"
	"gopkg.in/mgo.v2/bson"
)

type AccountDAO struct {
	AbstractDAO
}

func (accountDAO AccountDAO) Create(account *entities.Account) error {
	hash, _ := bcrypt.GenerateFromPassword([]byte(account.Password), bcrypt.DefaultCost)
	account.Password = string(hash)
	return accountDAO.DB.C(accountDAO.COLLECTION).Insert(&account)
}

func (accountDAO AccountDAO) CheckUsernameAndPassword(username, password string) bool {
	var account entities.Account
	err := accountDAO.DB.C(accountDAO.COLLECTION).Find(bson.M{
		"username": username,
	}).One(&account)
	if err != nil {
		return false
	} else {
		return bcrypt.CompareHashAndPassword([]byte(account.Password), []byte(password)) == nil
	}
}

func (accountDAO AccountDAO) CheckEmailAndPassword(email, password string) bool {
	var account entities.Account
	err := accountDAO.DB.C(accountDAO.COLLECTION).Find(bson.M{
		"email": email,
	}).One(&account)
	if err != nil {
		return false
	} else {
		return bcrypt.CompareHashAndPassword([]byte(account.Password), []byte(password)) == nil
	}
}
