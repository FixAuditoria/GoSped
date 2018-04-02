package Model

import (
	"gopkg.in/mgo.v2/bson"
)

// RegH005 : Totais do Inventário
type RegH005 struct {
	ID       bson.ObjectId `bson:"_id,omitempty" json:"id,omitempty"`
	Reg      string        `bson:"reg" json:"reg"`
	DtInv    string        `bson:"dtinv" json:"dtinv"`
	VlInv    string        `bson:"vlinv" json:"vlinv"`
	MotInv   string        `bson:"motinv" json:"motinv"`
	DtIni    string        `bson:"dtini" json:"dtini"`
	DtFin    string        `bson:"dtfin" json:"dtfin"`
	CnpjSped string        `bson:"cnpjsped" json:"cnpjsped"`
}

// Populate : O métdodo é responsável por preencher os dados pelo sped
func (r *RegH005) Populate(l []string, reg0000 Reg0000) {
	r.Reg = l[1]
	r.DtInv = l[2]
	r.VlInv = l[3]
	r.MotInv = l[4]
	r.DtIni = reg0000.DtIni
	r.DtFin = reg0000.DtFin
	r.CnpjSped = reg0000.Cnpj
}
