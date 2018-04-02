package Dao

import (
	"log"

	"github.com/chapzin/GoSped/Model"
	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type RegC405Dao struct {
	Server   string
	Database string
}

func (r *RegC405Dao) Connect() {
	session, err := mgo.Dial(r.Server)
	if err != nil {
		log.Fatal(err)
	}
	db = session.DB(r.Database)
}

// Lista todos os registro 0000
func (r *RegC405Dao) FindAll() ([]Model.RegC405, error) {
	var registros []Model.RegC405
	err := db.C(COLLECTIONC405).Find(bson.M{}).All(&registros)
	return registros, err
}

func (r *RegC405Dao) FindByCnpj(cnpjsped string) ([]Model.RegC405, error) {
	var registros []Model.RegC405
	err := db.C(COLLECTIONC405).Find(bson.M{"cnpjsped": cnpjsped}).All(&registros)
	return registros, err
}

func (r *RegC405Dao) FindById(id string) (Model.RegC405, error) {
	var regC405 Model.RegC405
	err := db.C(COLLECTIONC405).Find(bson.ObjectIdHex(id)).One(&regC405)
	return regC405, err
}

func (r *RegC405Dao) Insert(regC405 Model.RegC405) error {
	err := db.C(COLLECTIONC405).Insert(&regC405)
	return err
}

func (r *RegC405Dao) Delete(regC405 Model.RegC405) error {
	err := db.C(COLLECTIONC405).Remove(&regC405)
	return err
}

func (r *RegC405Dao) Update(regC405 Model.RegC405) error {
	err := db.C(COLLECTIONC405).UpdateId(regC405.ID, &regC405)
	return err
}
