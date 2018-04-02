package Model

import (
	"gopkg.in/mgo.v2/bson"
)

// RegH010 : Inventário
type RegH010 struct {
	ID       bson.ObjectId `bson:"_id,omitempty" json:"id,omitempty"`
	Reg      string        `bson:"reg" json:"reg"`
	CodItem  string        `bson:"coditem" json:"coditem"`
	Unid     string        `bson:"unid" json:"unid"`
	Qtd      string        `bson:"qtd" json:"qtd"`
	VlUnit   string        `bson:"vlunit" json:"vlunit"`
	VlItem   string        `bson:"vlitem" json:"vlitem"`
	IndProp  string        `bson:"indprop" json:"indprop"`
	CodPart  string        `bson:"codpart" json:"codpart"`
	TxtCompl string        `bson:"txtcompl" json:"txtcompl"`
	CodCta   string        `bson:"codcta" json:"codcta"`
	VlItemIr string        `bson:"vlitemir" json:"vlitemir"`
	DtInv    string        `bson:"dtinv" json:"dtinv"`
	DtIni    string        `bson:"dtini" json:"dtini"`
	DtFin    string        `bson:"dtfin" json:"dtfin"`
	CnpjSped string        `bson:"cnpjsped" json:"cnpjsped"`
}

// Populate : O métdodo é responsável por preencher os dados pelo sped
func (r *RegH010) Populate(l []string, reg0000 Reg0000, regH005 RegH005) {
	r.Reg = l[1]
	r.CodItem = l[2]
	r.Unid = l[3]
	r.Qtd = l[4]
	r.VlUnit = l[5]
	r.VlItem = l[6]
	r.IndProp = l[7]
	r.CodPart = l[8]
	r.TxtCompl = l[9]
	r.CodCta = l[10]
	r.VlItemIr = l[11]
	r.DtInv = regH005.DtInv
	r.DtIni = reg0000.DtIni
	r.DtFin = reg0000.DtFin
	r.CnpjSped = reg0000.Cnpj
}
