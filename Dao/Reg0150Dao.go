package Dao

import (
	"log"

	"github.com/chapzin/GoSped/Model"
	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type Reg0150Dao struct {
	Server   string
	Database string
}

func (r *Reg0150Dao) Connect() {
	session, err := mgo.Dial(r.Server)
	if err != nil {
		log.Fatal(err)
	}
	db = session.DB(r.Database)
}

// Lista todos os registro 0000
func (r *Reg0150Dao) FindAll() ([]Model.Reg0150, error) {
	var registros []Model.Reg0150
	err := db.C(COLLECTION0150).Find(bson.M{}).All(&registros)
	return registros, err
}

func (r *Reg0150Dao) FindById(id string) (Model.Reg0150, error) {
	var reg0150 Model.Reg0150
	err := db.C(COLLECTION0150).Find(bson.ObjectIdHex(id)).One(&reg0150)
	return reg0150, err
}

func (r *Reg0150Dao) Insert(reg0150 Model.Reg0150) error {
	err := db.C(COLLECTION0150).Insert(&reg0150)
	return err
}

func (r *Reg0150Dao) Delete(reg0150 Model.Reg0150) error {
	err := db.C(COLLECTION0150).Remove(&reg0150)
	return err
}

func (r *Reg0150Dao) Update(reg0150 Model.Reg0150) error {
	err := db.C(COLLECTION0150).UpdateId(reg0150.ID, &reg0150)
	return err
}
