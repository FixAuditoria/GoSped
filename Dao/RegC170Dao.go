package Dao

import (
	"log"

	"github.com/chapzin/GoSped/Model"
	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type RegC170Dao struct {
	Server   string
	Database string
}

func (r *RegC170Dao) Connect() {
	session, err := mgo.Dial(r.Server)
	if err != nil {
		log.Fatal(err)
	}
	db = session.DB(r.Database)
}

// Lista todos os registro 0000
func (r *RegC170Dao) FindAll() ([]Model.RegC170, error) {
	var registros []Model.RegC170
	err := db.C(COLLECTIONC170).Find(bson.M{}).All(&registros)
	return registros, err
}

func (r *RegC170Dao) FindById(id string) (Model.RegC170, error) {
	var regC170 Model.RegC170
	err := db.C(COLLECTIONC170).Find(bson.ObjectIdHex(id)).One(&regC170)
	return regC170, err
}

func (r *RegC170Dao) Insert(regC170 Model.RegC170) error {
	err := db.C(COLLECTIONC170).Insert(&regC170)
	return err
}

func (r *RegC170Dao) Delete(regC170 Model.RegC170) error {
	err := db.C(COLLECTIONC170).Remove(&regC170)
	return err
}

func (r *RegC170Dao) Update(regC170 Model.RegC170) error {
	err := db.C(COLLECTIONC170).UpdateId(regC170.ID, &regC170)
	return err
}
