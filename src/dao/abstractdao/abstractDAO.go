package dao

import (
	"../../entities"
	// . "../idao"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type AbstractDAO struct {
	// IDAO
	DB         *mgo.Database
	COLLECTION string
}

func (dao AbstractDAO) Create(entity entities.IEntity) error {
	return dao.DB.C(dao.COLLECTION).Insert(&entity)
}

func (dao AbstractDAO) FindAll() ([]map[string]interface{}, error) {
	var entities []map[string]interface{}
	err := dao.DB.C(dao.COLLECTION).Find(bson.M{}).All(&entities)
	if err != nil {
		return nil, err
	} else {
		return entities, nil
	}
}

func (dao AbstractDAO) Find(id string) (entities.IEntity, error) {
	var entity entities.IEntity
	err := dao.DB.C(dao.COLLECTION).FindId(bson.ObjectIdHex(id)).One(&entity)
	return entity, err
}

func (dao AbstractDAO) Delete(id string) error {
	return dao.DB.C(dao.COLLECTION).RemoveId(bson.ObjectIdHex(id))
}

func (dao AbstractDAO) Update(id string, entity entities.IEntity) error {
	return dao.DB.C(dao.COLLECTION).UpdateId(bson.ObjectIdHex(id), entity)
}
