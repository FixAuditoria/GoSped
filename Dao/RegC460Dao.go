package Dao

import (
	"log"

	"github.com/chapzin/GoSped/Model"
	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type RegC460Dao struct {
	Server   string
	Database string
}

func (r *RegC460Dao) Connect() {
	session, err := mgo.Dial(r.Server)
	if err != nil {
		log.Fatal(err)
	}
	db = session.DB(r.Database)
}

// Lista todos os registro 0000
func (r *RegC460Dao) FindAll() ([]Model.RegC460, error) {
	var registros []Model.RegC460
	err := db.C(COLLECTIONC460).Find(bson.M{}).All(&registros)
	return registros, err
}

func (r *RegC460Dao) FindByCnpj(cnpjsped string) ([]Model.RegC460, error) {
	var registros []Model.RegC460
	err := db.C(COLLECTIONC460).Find(bson.M{"cnpjsped": cnpjsped}).All(&registros)
	return registros, err
}

func (r *RegC460Dao) FindById(id string) (Model.RegC460, error) {
	var regC460 Model.RegC460
	err := db.C(COLLECTIONC460).Find(bson.ObjectIdHex(id)).One(&regC460)
	return regC460, err
}

func (r *RegC460Dao) Insert(regC460 Model.RegC460) error {
	err := db.C(COLLECTIONC460).Insert(&regC460)
	return err
}

func (r *RegC460Dao) Delete(regC460 Model.RegC460) error {
	err := db.C(COLLECTIONC460).Remove(&regC460)
	return err
}

func (r *RegC460Dao) Update(regC460 Model.RegC460) error {
	err := db.C(COLLECTIONC460).UpdateId(regC460.ID, &regC460)
	return err
}
