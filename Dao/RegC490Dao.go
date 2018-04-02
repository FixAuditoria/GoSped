package Dao

import (
	"log"

	"github.com/chapzin/GoSped/Model"
	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type RegC490Dao struct {
	Server   string
	Database string
}

func (r *RegC490Dao) Connect() {
	session, err := mgo.Dial(r.Server)
	if err != nil {
		log.Fatal(err)
	}
	db = session.DB(r.Database)
}

// Lista todos os registro 0000
func (r *RegC490Dao) FindAll() ([]Model.RegC490, error) {
	var registros []Model.RegC490
	err := db.C(COLLECTIONC490).Find(bson.M{}).All(&registros)
	return registros, err
}

func (r *RegC490Dao) FindByCnpj(cnpjsped string) ([]Model.RegC490, error) {
	var registros []Model.RegC490
	err := db.C(COLLECTIONC490).Find(bson.M{"cnpjsped": cnpjsped}).All(&registros)
	return registros, err
}

func (r *RegC490Dao) FindById(id string) (Model.RegC490, error) {
	var regC490 Model.RegC490
	err := db.C(COLLECTIONC490).Find(bson.ObjectIdHex(id)).One(&regC490)
	return regC490, err
}

func (r *RegC490Dao) Insert(regC490 Model.RegC490) error {
	err := db.C(COLLECTIONC490).Insert(&regC490)
	return err
}

func (r *RegC490Dao) Delete(regC490 Model.RegC490) error {
	err := db.C(COLLECTIONC490).Remove(&regC490)
	return err
}

func (r *RegC490Dao) Update(regC490 Model.RegC490) error {
	err := db.C(COLLECTIONC490).UpdateId(regC490.ID, &regC490)
	return err
}
