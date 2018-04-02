package Dao

import (
	"log"

	"github.com/chapzin/GoSped/Model"
	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type RegC425Dao struct {
	Server   string
	Database string
}

func (r *RegC425Dao) Connect() {
	session, err := mgo.Dial(r.Server)
	if err != nil {
		log.Fatal(err)
	}
	db = session.DB(r.Database)
}

// Lista todos os registro 0000
func (r *RegC425Dao) FindAll() ([]Model.RegC425, error) {
	var registros []Model.RegC425
	err := db.C(COLLECTIONC425).Find(bson.M{}).All(&registros)
	return registros, err
}

func (r *RegC425Dao) FindByCnpj(cnpjsped string) ([]Model.RegC425, error) {
	var registros []Model.RegC425
	err := db.C(COLLECTIONC425).Find(bson.M{"cnpjsped": cnpjsped}).All(&registros)
	return registros, err
}

func (r *RegC425Dao) FindById(id string) (Model.RegC425, error) {
	var regC425 Model.RegC425
	err := db.C(COLLECTIONC425).Find(bson.ObjectIdHex(id)).One(&regC425)
	return regC425, err
}

func (r *RegC425Dao) Insert(regC425 Model.RegC425) error {
	err := db.C(COLLECTIONC425).Insert(&regC425)
	return err
}

func (r *RegC425Dao) Delete(regC425 Model.RegC425) error {
	err := db.C(COLLECTIONC425).Remove(&regC425)
	return err
}

func (r *RegC425Dao) Update(regC425 Model.RegC425) error {
	err := db.C(COLLECTIONC425).UpdateId(regC425.ID, &regC425)
	return err
}
