package Dao

import (
	"log"

	"github.com/chapzin/GoSped/Model"
	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type Reg0220Dao struct {
	Server   string
	Database string
}

func (r *Reg0220Dao) Connect() {
	session, err := mgo.Dial(r.Server)
	if err != nil {
		log.Fatal(err)
	}
	db = session.DB(r.Database)
}

// Lista todos os registro 0000
func (r *Reg0220Dao) FindAll() ([]Model.Reg0220, error) {
	var registros []Model.Reg0220
	err := db.C(COLLECTION0220).Find(bson.M{}).All(&registros)
	return registros, err
}

func (r *Reg0220Dao) FindById(id string) (Model.Reg0220, error) {
	var reg0220 Model.Reg0220
	err := db.C(COLLECTION0220).Find(bson.ObjectIdHex(id)).One(&reg0220)
	return reg0220, err
}

func (r *Reg0220Dao) Insert(reg0220 Model.Reg0220) error {
	err := db.C(COLLECTION0220).Insert(&reg0220)
	return err
}

func (r *Reg0220Dao) Delete(reg0220 Model.Reg0220) error {
	err := db.C(COLLECTION0220).Remove(&reg0220)
	return err
}

func (r *Reg0220Dao) Update(reg0220 Model.Reg0220) error {
	err := db.C(COLLECTION0220).UpdateId(reg0220.ID, &reg0220)
	return err
}
