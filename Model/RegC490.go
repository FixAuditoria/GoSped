package Model

import "github.com/go-bongo/bongo"

// RegC490 : Registro Analítico do movimento diário (código 02, 2D e 60)
type RegC490 struct {
	bongo.DocumentBase `bson:",inline"`
	Reg                string
	CstIcms            string
	Cfop               string
	AliqIcms           string
	VlOpr              string
	VlBcIcms           string
	VlIcms             string
	CodObs             string
	DtIni              string
	DtFin              string
	CnpjSped           string
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
