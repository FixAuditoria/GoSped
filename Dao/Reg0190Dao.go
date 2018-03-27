package Dao

import (
	"log"

	"github.com/chapzin/GoSped/Model"
	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type Reg0190Dao struct {
	Server   string
	Database string
}

func (r *Reg0190Dao) Connect() {
	session, err := mgo.Dial(r.Server)
	if err != nil {
		log.Fatal(err)
	}
	db = session.DB(r.Database)
}

// Lista todos os registro 0000
func (r *Reg0190Dao) FindAll() ([]Model.Reg0190, error) {
	var registros []Model.Reg0190
	err := db.C(COLLECTION0190).Find(bson.M{}).All(&registros)
	return registros, err
}

func (r *Reg0190Dao) FindById(id string) (Model.Reg0190, error) {
	var reg0190 Model.Reg0190
	err := db.C(COLLECTION0190).Find(bson.ObjectIdHex(id)).One(&reg0190)
	return reg0190, err
}

func (r *Reg0190Dao) Insert(reg0190 Model.Reg0190) error {
	err := db.C(COLLECTION0190).Insert(&reg0190)
	return err
}

func (r *Reg0190Dao) Delete(reg0190 Model.Reg0190) error {
	err := db.C(COLLECTION0190).Remove(&reg0190)
	return err
}

func (r *Reg0190Dao) Update(reg0190 Model.Reg0190) error {
	err := db.C(COLLECTION0190).UpdateId(reg0190.ID, &reg0190)
	return err
}
