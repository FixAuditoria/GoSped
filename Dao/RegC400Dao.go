package Dao

import (
	"log"

	"github.com/chapzin/GoSped/Model"
	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type RegC400Dao struct {
	Server   string
	Database string
}

func (r *RegC400Dao) Connect() {
	session, err := mgo.Dial(r.Server)
	if err != nil {
		log.Fatal(err)
	}
	db = session.DB(r.Database)
}

// Lista todos os registro 0000
func (r *RegC400Dao) FindAll() ([]Model.RegC400, error) {
	var registros []Model.RegC400
	err := db.C(COLLECTIONC400).Find(bson.M{}).All(&registros)
	return registros, err
}

func (r *RegC400Dao) FindByCnpj(cnpjsped string) ([]Model.RegC400, error) {
	var registros []Model.RegC400
	err := db.C(COLLECTIONC400).Find(bson.M{"cnpjsped": cnpjsped}).All(&registros)
	return registros, err
}

func (r *RegC400Dao) FindById(id string) (Model.RegC400, error) {
	var regC400 Model.RegC400
	err := db.C(COLLECTIONC400).Find(bson.ObjectIdHex(id)).One(&regC400)
	return regC400, err
}

func (r *RegC400Dao) Insert(regC400 Model.RegC400) error {
	err := db.C(COLLECTIONC400).Insert(&regC400)
	return err
}

func (r *RegC400Dao) Delete(regC400 Model.RegC400) error {
	err := db.C(COLLECTIONC400).Remove(&regC400)
	return err
}

func (r *RegC400Dao) Update(regC400 Model.RegC400) error {
	err := db.C(COLLECTIONC400).UpdateId(regC400.ID, &regC400)
	return err
}
