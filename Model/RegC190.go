package Model

import "github.com/go-bongo/bongo"

// RegC190 : Registro Analítico do Documento (código 01, 1B, 04, 55 e 65)
type RegC190 struct {
	bongo.DocumentBase `bson:",inline"`
	Reg                string
	CstIcms            string
	AliqIcms           string
	VlOpr              string
	VlBcIcms           string
	VlIcms             string
	VlBcIcmsSt         string
	VlIcmsSt           string
	VlRedBc            string
	VlIpi              string
	CodObs             string
	DtIni              string
	DtFin              string
	CnpjSped           string
}

// Populate : O métdodo é responsável por preencher os dados pelo sped
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
