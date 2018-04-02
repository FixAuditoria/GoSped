package Dao

import (
	"log"

	"github.com/chapzin/GoSped/Model"
	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type ProdutosNfeDao struct {
	Server   string
	Database string
}

func (r *ProdutosNfeDao) Connect() {
	session, err := mgo.Dial(r.Server)
	if err != nil {
		log.Fatal(err)
	}
	db = session.DB(r.Database)
}

// Lista todos os registro 0000
func (r *ProdutosNfeDao) FindAll() ([]Model.ProdutosNfe, error) {
	var registros []Model.ProdutosNfe
	err := db.C(COLLECTIONDADOSPRODUTOS).Find(bson.M{}).All(&registros)
	return registros, err
}

func (r *ProdutosNfeDao) FindById(id string) (Model.ProdutosNfe, error) {
	var ProdutosNfe Model.ProdutosNfe
	err := db.C(COLLECTIONDADOSPRODUTOS).Find(bson.ObjectIdHex(id)).One(&ProdutosNfe)
	return ProdutosNfe, err
}

func (r *ProdutosNfeDao) Insert(ProdutosNfe Model.ProdutosNfe) error {
	err := db.C(COLLECTIONDADOSPRODUTOS).Insert(&ProdutosNfe)
	return err
}

func (r *ProdutosNfeDao) Delete(ProdutosNfe Model.ProdutosNfe) error {
	err := db.C(COLLECTIONDADOSPRODUTOS).Remove(&ProdutosNfe)
	return err
}

func (r *ProdutosNfeDao) Update(ProdutosNfe Model.ProdutosNfe) error {
	err := db.C(COLLECTIONDADOSPRODUTOS).UpdateId(ProdutosNfe.ID, &ProdutosNfe)
	return err
}
