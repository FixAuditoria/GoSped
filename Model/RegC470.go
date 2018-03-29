package Model

import (
	"gopkg.in/mgo.v2/bson"
)

// RegC470 : Itens do Documento Fiscal Emitido por ECF (código 02 e 2D)
type RegC470 struct {
	ID       bson.ObjectId `bson:"_id,omitempty" json:"id,omitempty"`
	Reg      string        `bson:"reg" json:"reg"`
	CodItem  string        `bson:"coditem" json:"coditem"`
	Qtd      string        `bson:"qtd" json:"qtd"`
	QtdCanc  string        `bson:"qtdcanc" json:"qtdcanc"`
	Unid     string        `bson:"unid" json:"unid"`
	VlItem   string        `bson:"vlitem" json:"vlitem"`
	CstIcms  string        `bson:"csticms" json:"csticms"`
	Cfop     string        `bson:"cfop" json:"cfop"`
	AliqIcms string        `bson:"aliqicms" json:"aliqicms"`
	VlPis    string        `bson:"vlpis" json:"vlpis"`
	VlCofins string        `bson:"vlcofins" json:"vlcofins"`
	DtIni    string        `bson:"dtini" json:"dtini"`
	DtFin    string        `bson:"dtfin" json:"dtfin"`
	CnpjSped string        `bson:"cnpjsped" json:"cnpjsped"`
}

// Populate: O métdodo é responsável por preencher os dados pelo sped
func (r *RegC470) Populate(l []string, reg0000 Reg0000) {
	r.Reg = l[1]
	r.CodItem = l[2]
	r.Qtd = l[3]
	r.QtdCanc = l[4]
	r.Unid = l[5]
	r.VlItem = l[6]
	r.CstIcms = l[7]
	r.Cfop = l[8]
	r.AliqIcms = l[9]
	r.VlPis = l[10]
	r.VlCofins = l[11]
	r.DtIni = reg0000.DtIni
	r.DtFin = reg0000.DtFin
	r.CnpjSped = reg0000.Cnpj
}
