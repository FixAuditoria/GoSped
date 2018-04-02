package Dao

import (
	"log"

	"github.com/chapzin/GoSped/Model"
	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type RegH010Dao struct {
	Server   string
	Database string
}

func (r *RegH010Dao) Connect() {
	session, err := mgo.Dial(r.Server)
	if err != nil {
		log.Fatal(err)
	}
	db = session.DB(r.Database)
}

// Lista todos os registro 0000
func (r *RegH010Dao) FindAll() ([]Model.RegH010, error) {
	var registros []Model.RegH010
	err := db.C(COLLECTIONH010).Find(bson.M{}).All(&registros)
	return registros, err
}

func (r *RegH010Dao) FindByCnpj(cnpjsped string) ([]Model.RegH010, error) {
	var registros []Model.RegH010
	err := db.C(COLLECTIONH010).Find(bson.M{"cnpjsped": cnpjsped}).All(&registros)
	return registros, err
}

func (r *RegH010Dao) FindById(id string) (Model.RegH010, error) {
	var regH010 Model.RegH010
	err := db.C(COLLECTIONH010).Find(bson.ObjectIdHex(id)).One(&regH010)
	return regH010, err
}

func (r *RegH010Dao) Insert(regH010 Model.RegH010) error {
	err := db.C(COLLECTIONH010).Insert(&regH010)
	return err
}

func (r *RegH010Dao) Delete(regH010 Model.RegH010) error {
	err := db.C(COLLECTIONH010).Remove(&regH010)
	return err
}

func (r *RegH010Dao) Update(regH010 Model.RegH010) error {
	err := db.C(COLLECTIONH010).UpdateId(regH010.ID, &regH010)
	return err
}
