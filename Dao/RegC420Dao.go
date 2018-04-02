package Dao

import (
	"log"

	"github.com/chapzin/GoSped/Model"
	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type RegC420Dao struct {
	Server   string
	Database string
}

func (r *RegC420Dao) Connect() {
	session, err := mgo.Dial(r.Server)
	if err != nil {
		log.Fatal(err)
	}
	db = session.DB(r.Database)
}

// Lista todos os registro 0000
func (r *RegC420Dao) FindAll() ([]Model.RegC420, error) {
	var registros []Model.RegC420
	err := db.C(COLLECTIONC420).Find(bson.M{}).All(&registros)
	return registros, err
}

func (r *RegC420Dao) FindByCnpj(cnpjsped string) ([]Model.RegC420, error) {
	var registros []Model.RegC420
	err := db.C(COLLECTIONC420).Find(bson.M{"cnpjsped": cnpjsped}).All(&registros)
	return registros, err
}

func (r *RegC420Dao) FindById(id string) (Model.RegC420, error) {
	var regC420 Model.RegC420
	err := db.C(COLLECTIONC420).Find(bson.ObjectIdHex(id)).One(&regC420)
	return regC420, err
}

func (r *RegC420Dao) Insert(regC420 Model.RegC420) error {
	err := db.C(COLLECTIONC420).Insert(&regC420)
	return err
}

func (r *RegC420Dao) Delete(regC420 Model.RegC420) error {
	err := db.C(COLLECTIONC420).Remove(&regC420)
	return err
}

func (r *RegC420Dao) Update(regC420 Model.RegC420) error {
	err := db.C(COLLECTIONC420).UpdateId(regC420.ID, &regC420)
	return err
}
