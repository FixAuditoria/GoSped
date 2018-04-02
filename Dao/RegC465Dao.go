package Dao

import (
	"log"

	"github.com/chapzin/GoSped/Model"
	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type RegC465Dao struct {
	Server   string
	Database string
}

func (r *RegC465Dao) Connect() {
	session, err := mgo.Dial(r.Server)
	if err != nil {
		log.Fatal(err)
	}
	db = session.DB(r.Database)
}

// Lista todos os registro 0000
func (r *RegC465Dao) FindAll() ([]Model.RegC465, error) {
	var registros []Model.RegC465
	err := db.C(COLLECTIONC465).Find(bson.M{}).All(&registros)
	return registros, err
}

func (r *RegC465Dao) FindByCnpj(cnpjsped string) ([]Model.RegC465, error) {
	var registros []Model.RegC465
	err := db.C(COLLECTIONC465).Find(bson.M{"cnpjsped": cnpjsped}).All(&registros)
	return registros, err
}

func (r *RegC465Dao) FindById(id string) (Model.RegC465, error) {
	var regC465 Model.RegC465
	err := db.C(COLLECTIONC465).Find(bson.ObjectIdHex(id)).One(&regC465)
	return regC465, err
}

func (r *RegC465Dao) Insert(regC465 Model.RegC465) error {
	err := db.C(COLLECTIONC465).Insert(&regC465)
	return err
}

func (r *RegC465Dao) Delete(regC465 Model.RegC465) error {
	err := db.C(COLLECTIONC465).Remove(&regC465)
	return err
}

func (r *RegC465Dao) Update(regC465 Model.RegC465) error {
	err := db.C(COLLECTIONC465).UpdateId(regC465.ID, &regC465)
	return err
}
