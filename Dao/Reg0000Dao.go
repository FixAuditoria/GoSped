package Dao

import (
	"log"

	"github.com/chapzin/GoSped/Model"
	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type Reg0000Dao struct {
	Server   string
	Database string
}

func (r *Reg0000Dao) Connect() {
	session, err := mgo.Dial(r.Server)
	if err != nil {
		log.Fatal(err)
	}
	db = session.DB(r.Database)
}

// Lista todos os registro 0000
func (r *Reg0000Dao) FindAll() ([]Model.Reg0000, error) {
	var registros []Model.Reg0000
	err := db.C(COLLECTION0000).Find(bson.M{}).All(&registros)
	return registros, err
}

func (r *Reg0000Dao) FindById(id string) (Model.Reg0000, error) {
	var reg0000 Model.Reg0000
	err := db.C(COLLECTION0000).Find(bson.ObjectIdHex(id)).One(&reg0000)
	return reg0000, err
}

func (r *Reg0000Dao) Insert(reg0000 Model.Reg0000) error {
	err := db.C(COLLECTION0000).Insert(&reg0000)
	return err
}

func (r *Reg0000Dao) Delete(reg0000 Model.Reg0000) error {
	err := db.C(COLLECTION0000).Remove(&reg0000)
	return err
}

func (r *Reg0000Dao) Update(reg0000 Model.Reg0000) error {
	err := db.C(COLLECTION0000).UpdateId(reg0000.ID, &reg0000)
	return err
}
