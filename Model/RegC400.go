package Model

import "gopkg.in/mgo.v2/bson"

// RegC400 : Equipamento ECF (código 02, 2D e 60),
type RegC400 struct {
	ID       bson.ObjectId `bson:"_id,omitempty" json:"id,omitempty"`
	Reg      string        `bson:"reg" json:"reg"`
	CodMod   string        `bson:"codmod" json:"codmod"`
	EcfMod   string        `bson:"ecfmod" json:"ecfmod"`
	EcfFab   string        `bson:"ecffab" json:"ecffab"`
	EcfCx    string        `bson:"ecfcx" json:"ecfcx"`
	DtIni    string        `bson:"dtini" json:"dtini"`
	DtFin    string        `bson:"dtfin" json:"dtfin"`
	CnpjSped string        `bson:"cnpjsped" json:"cnpjsped"`
}

// Populate : O métdodo é responsável por preencher os dados pelo sped
func (r *RegC400) Populate(l []string, reg0000 Reg0000) {
	r.Reg = l[1]
	r.CodMod = l[2]
	r.EcfMod = l[3]
	r.EcfFab = l[4]
	r.EcfCx = l[5]
	r.DtIni = reg0000.DtIni
	r.DtFin = reg0000.DtFin
	r.CnpjSped = reg0000.Cnpj
}
