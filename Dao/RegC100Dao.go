package Dao

import (
	"log"

	"github.com/chapzin/GoSped/Model"
	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type RegC100Dao struct {
	Server   string
	Database string
}

func (r *RegC100Dao) Connect() {
	session, err := mgo.Dial(r.Server)
	if err != nil {
		log.Fatal(err)
	}
	db = session.DB(r.Database)
}

// Lista todos os registro 0000
func (r *RegC100Dao) FindAll() ([]Model.RegC100, error) {
	var registros []Model.RegC100
	err := db.C(COLLECTIONC100).Find(bson.M{}).All(&registros)
	return registros, err
}

func (r *RegC100Dao) FindById(id string) (Model.RegC100, error) {
	var regC100 Model.RegC100
	err := db.C(COLLECTIONC100).Find(bson.ObjectIdHex(id)).One(&regC100)
	return regC100, err
}

func (r *RegC100Dao) Insert(regC100 Model.RegC100) error {
	err := db.C(COLLECTIONC100).Insert(&regC100)
	return err
}

func (r *RegC100Dao) Delete(regC100 Model.RegC100) error {
	err := db.C(COLLECTIONC100).Remove(&regC100)
	return err
}

func (r *RegC100Dao) Update(regC100 Model.RegC100) error {
	err := db.C(COLLECTIONC100).UpdateId(regC100.ID, &regC100)
	return err
}
