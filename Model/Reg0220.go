package Model

import (
	"gopkg.in/mgo.v2/bson"
)

// Reg0220 : Fatores de Conversão de Unidades
type Reg0220 struct {
	ID       bson.ObjectId `bson:"_id,omitempty" json:"id,omitempty"`
	Reg      string        `bson:"reg" json:"reg"`
	UnidConv string        `bson:"unidconv" json:"unidconv"`
	FatConv  string        `bson:"fatconv" json:"fatconv"`
	UnidCod  string        `bson:"unidcod" json:"unidcod"`
	CodItem  string        `bson:"coditem" json:"coditem"`
	DtIni    string        `bson:"dtini" json:"dtini"`
	DtFin    string        `bson:"dtfin" json:"dtfin"`
	CnpjSped string        `bson:"cnpj" json:"cnpj"`
}

// Populate: O métdodo é responsável por preencher os dados pelo sped
func (r *Reg0220) Populate(l []string, reg0000 Reg0000, reg0200 Reg0200) {
	r.Reg = l[1]
	r.UnidConv = l[2]
	r.FatConv = l[3]
	r.UnidCod = reg0200.UnidInv
	r.CodItem = reg0200.CodItem
	r.DtIni = reg0000.DtIni
	r.DtFin = reg0000.DtFin
	r.CnpjSped = reg0000.Cnpj
}
