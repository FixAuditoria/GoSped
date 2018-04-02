package Dao

import (
	"log"

	"github.com/chapzin/GoSped/Model"
	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type RegH005Dao struct {
	Server   string
	Database string
}

func (r *RegH005Dao) Connect() {
	session, err := mgo.Dial(r.Server)
	if err != nil {
		log.Fatal(err)
	}
	db = session.DB(r.Database)
}

// Lista todos os registro 0000
func (r *RegH005Dao) FindAll() ([]Model.RegH005, error) {
	var registros []Model.RegH005
	err := db.C(COLLECTIONH005).Find(bson.M{}).All(&registros)
	return registros, err
}

func (r *RegH005Dao) FindByCnpj(cnpjsped string) ([]Model.RegH005, error) {
	var registros []Model.RegH005
	err := db.C(COLLECTIONH005).Find(bson.M{"cnpjsped": cnpjsped}).All(&registros)
	return registros, err
}

func (r *RegH005Dao) FindById(id string) (Model.RegH005, error) {
	var regH005 Model.RegH005
	err := db.C(COLLECTIONH005).Find(bson.ObjectIdHex(id)).One(&regH005)
	return regH005, err
}

func (r *RegH005Dao) Insert(regH005 Model.RegH005) error {
	err := db.C(COLLECTIONH005).Insert(&regH005)
	return err
}

func (r *RegH005Dao) Delete(regH005 Model.RegH005) error {
	err := db.C(COLLECTIONH005).Remove(&regH005)
	return err
}

func (r *RegH005Dao) Update(regH005 Model.RegH005) error {
	err := db.C(COLLECTIONH005).UpdateId(regH005.ID, &regH005)
	return err
}
