package Model

import (
	"gopkg.in/mgo.v2/bson"
)

// RegC190 : Registro Analítico do Documento (código 01, 1B, 04, 55 e 65)
type RegC190 struct {
	ID         bson.ObjectId `bson:"_id,omitempty" json:"id,omitempty"`
	Reg        string        `bson:"reg" json:"reg"`
	CstIcms    string        `bson:"csticms" json:"csticms"`
	AliqIcms   string        `bson:"aliqicms" json:"aliqicms"`
	VlOpr      string        `bson:"vlopr" json:"vlopr"`
	VlBcIcms   string        `bson:"vlbcicms" json:"vlbcicms"`
	VlIcms     string        `bson:"vlicms" json:"vlicms"`
	VlBcIcmsSt string        `bson:"vlbcicmsst" json:"vlbcicmsst"`
	VlIcmsSt   string        `bson:"vlicmsst" json:"vlicmsst"`
	VlRedBc    string        `bson:"vlredbc" json:"vlredbc"`
	VlIpi      string        `bson:"vlipi" json:"vlipi"`
	CodObs     string        `bson:"codobs" json:"codobs"`
	DtIni      string        `bson:"dtini" json:"dtini"`
	DtFin      string        `bson:"dtfin" json:"dtfin"`
	CnpjSped   string        `bson:"cnpjsped" json:"cnpjsped"`
}

// Populate: O métdodo é responsável por preencher os dados pelo sped
func (r *RegC190) Populate(l []string, reg0000 Reg0000) {
	r.Reg = l[1]
	r.CstIcms = l[2]
	r.AliqIcms = l[3]
	r.VlOpr = l[4]
	r.VlBcIcms = l[5]
	r.VlIcms = l[6]
	r.VlBcIcmsSt = l[7]
	r.VlIcmsSt = l[8]
	r.VlRedBc = l[9]
	r.VlIpi = l[10]
	r.CodObs = l[11]
	r.DtIni = reg0000.DtIni
	r.DtFin = reg0000.DtFin
	r.CnpjSped = reg0000.Cnpj
}
