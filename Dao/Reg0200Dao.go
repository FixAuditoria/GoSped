package Dao

import (
	"log"

	"github.com/chapzin/GoSped/Model"
	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type Reg0200Dao struct {
	Server   string
	Database string
}

func (r *Reg0200Dao) Connect() {
	session, err := mgo.Dial(r.Server)
	if err != nil {
		log.Fatal(err)
	}
	db = session.DB(r.Database)
}

// Lista todos os registro 0000
func (r *Reg0200Dao) FindAll() ([]Model.Reg0200, error) {
	var registros []Model.Reg0200
	err := db.C(COLLECTION0200).Find(bson.M{}).All(&registros)
	return registros, err
}

func (r *Reg0200Dao) FindById(id string) (Model.Reg0200, error) {
	var reg0200 Model.Reg0200
	err := db.C(COLLECTION0200).Find(bson.ObjectIdHex(id)).One(&reg0200)
	return reg0200, err
}

func (r *Reg0200Dao) Insert(reg0200 Model.Reg0200) error {
	err := db.C(COLLECTION0200).Insert(&reg0200)
	return err
}

func (r *Reg0200Dao) Delete(reg0200 Model.Reg0200) error {
	err := db.C(COLLECTION0200).Remove(&reg0200)
	return err
}

func (r *Reg0200Dao) Update(reg0200 Model.Reg0200) error {
	err := db.C(COLLECTION0200).UpdateId(reg0200.ID, &reg0200)
	return err
}
