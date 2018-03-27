package Dao

import (
	"log"

	"github.com/chapzin/GoSped/Model"
	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type Reg0100Dao struct {
	Server   string
	Database string
}

func (r *Reg0100Dao) Connect() {
	session, err := mgo.Dial(r.Server)
	if err != nil {
		log.Fatal(err)
	}
	db = session.DB(r.Database)
}

// Lista todos os registro 0100
func (r *Reg0100Dao) FindAll() ([]Model.Reg0100, error) {
	var registros []Model.Reg0100
	err := db.C(COLLECTION0100).Find(bson.M{}).All(&registros)
	return registros, err
}

func (r *Reg0100Dao) FindById(id string) (Model.Reg0100, error) {
	var reg0100 Model.Reg0100
	err := db.C(COLLECTION0100).Find(bson.ObjectIdHex(id)).One(&reg0100)
	return reg0100, err
}

func (r *Reg0100Dao) Insert(reg0100 Model.Reg0100) error {
	err := db.C(COLLECTION0100).Insert(&reg0100)
	return err
}

func (r *Reg0100Dao) Delete(reg0100 Model.Reg0100) error {
	err := db.C(COLLECTION0100).Remove(&reg0100)
	return err
}

func (r *Reg0100Dao) Update(reg0100 Model.Reg0100) error {
	err := db.C(COLLECTION0100).UpdateId(reg0100.ID, &reg0100)
	return err
}
