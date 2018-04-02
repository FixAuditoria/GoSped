package Dao

import (
	"log"

	"github.com/chapzin/GoSped/Model"
	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type RegC470Dao struct {
	Server   string
	Database string
}

func (r *RegC470Dao) Connect() {
	session, err := mgo.Dial(r.Server)
	if err != nil {
		log.Fatal(err)
	}
	db = session.DB(r.Database)
}

// Lista todos os registro 0000
func (r *RegC470Dao) FindAll() ([]Model.RegC470, error) {
	var registros []Model.RegC470
	err := db.C(COLLECTIONC470).Find(bson.M{}).All(&registros)
	return registros, err
}

func (r *RegC470Dao) FindByCnpj(cnpjsped string) ([]Model.RegC470, error) {
	var registros []Model.RegC470
	err := db.C(COLLECTIONC470).Find(bson.M{"cnpjsped": cnpjsped}).All(&registros)
	return registros, err
}

func (r *RegC470Dao) FindById(id string) (Model.RegC470, error) {
	var regC470 Model.RegC470
	err := db.C(COLLECTIONC470).Find(bson.ObjectIdHex(id)).One(&regC470)
	return regC470, err
}

func (r *RegC470Dao) Insert(regC470 Model.RegC470) error {
	err := db.C(COLLECTIONC470).Insert(&regC470)
	return err
}

func (r *RegC470Dao) Delete(regC470 Model.RegC470) error {
	err := db.C(COLLECTIONC470).Remove(&regC470)
	return err
}

func (r *RegC470Dao) Update(regC470 Model.RegC470) error {
	err := db.C(COLLECTIONC470).UpdateId(regC470.ID, &regC470)
	return err
}
