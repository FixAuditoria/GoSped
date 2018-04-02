package Dao

import (
	"log"

	"github.com/chapzin/GoSped/Model"
	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type RegC190Dao struct {
	Server   string
	Database string
}

func (r *RegC190Dao) Connect() {
	session, err := mgo.Dial(r.Server)
	if err != nil {
		log.Fatal(err)
	}
	db = session.DB(r.Database)
}

// Lista todos os registro 0000
func (r *RegC190Dao) FindAll() ([]Model.RegC190, error) {
	var registros []Model.RegC190
	err := db.C(COLLECTIONC190).Find(bson.M{}).All(&registros)
	return registros, err
}

func (r *RegC190Dao) FindByCnpj(cnpjsped string) ([]Model.RegC190, error) {
	var registros []Model.RegC190
	err := db.C(COLLECTIONC190).Find(bson.M{"cnpjsped": cnpjsped}).All(&registros)
	return registros, err
}

func (r *RegC190Dao) FindById(id string) (Model.RegC190, error) {
	var regC190 Model.RegC190
	err := db.C(COLLECTIONC190).Find(bson.ObjectIdHex(id)).One(&regC190)
	return regC190, err
}

func (r *RegC190Dao) Insert(regC190 Model.RegC190) error {
	err := db.C(COLLECTIONC190).Insert(&regC190)
	return err
}

func (r *RegC190Dao) Delete(regC190 Model.RegC190) error {
	err := db.C(COLLECTIONC190).Remove(&regC190)
	return err
}

func (r *RegC190Dao) Update(regC190 Model.RegC190) error {
	err := db.C(COLLECTIONC190).UpdateId(regC190.ID, &regC190)
	return err
}
