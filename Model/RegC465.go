package Model

import (
	"gopkg.in/mgo.v2/bson"
)

// RegC465 : Complemento do Cupom Fiscal Eletrônico Emitido por ECF - CF-e-ECF (código 60)
type RegC465 struct {
	ID       bson.ObjectId `bson:"_id,omitempty" json:"id,omitempty"`
	Reg      string        `bson:"reg" json:"reg"`
	ChvCfe   string        `bson:"chvcfe" json:"chvcfe"`
	NumCcF   string        `bson:"numccf" json:"numccf"`
	DtIni    string        `bson:"dtini" json:"dtini"`
	DtFin    string        `bson:"dtfin" json:"dtfin"`
	CnpjSped string        `bson:"cnpjsped" json:"cnpjsped"`
}

// Populate : O métdodo é responsável por preencher os dados pelo sped
func (r *RegC465) Populate(l []string, reg0000 Reg0000) {
	r.Reg = l[1]
	r.ChvCfe = l[2]
	r.NumCcF = l[3]
	r.DtIni = reg0000.DtIni
	r.DtFin = reg0000.DtFin
	r.CnpjSped = reg0000.Cnpj
}
