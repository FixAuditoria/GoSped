package Model

import (
	"gopkg.in/mgo.v2/bson"
)

// RegC490 : Registro Analítico do movimento diário (código 02, 2D e 60)
type RegC490 struct {
	ID       bson.ObjectId `bson:"_id,omitempty" json:"id,omitempty"`
	Reg      string        `bson:"reg" json:"reg"`
	CstIcms  string        `bson:"csticms" json:"csticms"`
	Cfop     string        `bson:"cfop" json:"cfop"`
	AliqIcms string        `bson:"aliqicms" json:"aliqicms"`
	VlOpr    string        `bson:"vlopr" json:"vlopr"`
	VlBcIcms string        `bson:"vlbcicms" json:"vlbcicms"`
	VlIcms   string        `bson:"vlicms" json:"vlicms"`
	CodObs   string        `bson:"codobs" json:"codobs"`
	DtIni    string        `bson:"dtini" json:"dtini"`
	DtFin    string        `bson:"dtfin" json:"dtfin"`
	CnpjSped string        `bson:"cnpjsped" json:"cnpjsped"`
}

// Populate : O métdodo é responsável por preencher os dados pelo sped
func (r *RegC490) Populate(l []string, reg0000 Reg0000) {
	r.Reg = l[1]
	r.CstIcms = l[2]
	r.Cfop = l[3]
	r.AliqIcms = l[4]
	r.VlOpr = l[5]
	r.VlBcIcms = l[6]
	r.VlIcms = l[7]
	r.CodObs = l[8]
	r.DtIni = reg0000.DtIni
	r.DtFin = reg0000.DtFin
	r.CnpjSped = reg0000.Cnpj
}
