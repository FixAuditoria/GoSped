package Model

import (
	"gopkg.in/mgo.v2/bson"
)

// RegC460 : Documento Fiscal Emitido por ECF (código 02, 2D e 60)
type RegC460 struct {
	ID       bson.ObjectId `bson:"_id,omitempty" json:"id,omitempty"`
	Reg      string        `bson:"reg" json:"reg"`
	CodMod   string        `bson:"codmod" json:"codmod"`
	CodSit   string        `bson:"codsit" json:"codsit"`
	NumDoc   string        `bson:"numdoc" json:"numdoc"`
	DtDoc    string        `bson:"dtdoc" json:"dtdoc"`
	VlDoc    string        `bson:"vldoc" json:"vldoc"`
	VlPis    string        `bson:"vlpis" json:"vlpis"`
	VlCofins string        `bson:"vlcofins" json:"vlcofins"`
	CpfCnpj  string        `bson:"cpfcnpj" json:"cpfcnpj"`
	NomAdq   string        `bson:"nomadq" json:"nomadq"`
	DtIni    string        `bson:"dtini" json:"dtini"`
	DtFin    string        `bson:"dtfin" json:"dtfin"`
	CnpjSped string        `bson:"cnpjsped" json:"cnpjsped"`
}

// Populate : O métdodo é responsável por preencher os dados pelo sped
func (r *RegC460) Populate(l []string, reg0000 Reg0000) {
	r.Reg = l[1]
	r.CodMod = l[2]
	r.CodSit = l[3]
	r.NumDoc = l[4]
	r.DtDoc = l[5]
	r.VlDoc = l[6]
	r.VlPis = l[7]
	r.VlCofins = l[8]
	r.CpfCnpj = l[9]
	r.NomAdq = l[10]
	r.DtIni = reg0000.DtIni
	r.DtFin = reg0000.DtFin
	r.CnpjSped = reg0000.Cnpj
}
