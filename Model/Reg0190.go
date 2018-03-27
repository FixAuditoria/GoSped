package Model

import (
	"gopkg.in/mgo.v2/bson"
)

// Reg0190 : Identificação das unidades de medida
type Reg0190 struct {
	ID       bson.ObjectId `bson:"_id,omitempty" json:"id,omitempty"`
	Reg      string        `bson:"reg" json:"reg"`
	Unid     string        `bson:"unid" json:"unid"`
	Descr    string        `bson:"descr" json:"descr"`
	DtIni    string        `bson:"dtini" json:"dtini"`
	DtFin    string        `bson:"dtfin" json:"dtfin"`
	CnpjSped string        `bson:"cnpj" json:"cnpj"`
}

// Populate: O métdodo é responsável por preencher os dados pelo sped
func (r *Reg0190) Populate(l []string, reg0000 Reg0000) {
	r.Reg = l[1]
	r.Unid = l[2]
	r.Descr = l[3]
	r.DtIni = reg0000.DtIni
	r.DtFin = reg0000.DtFin
	r.CnpjSped = reg0000.Cnpj
}
