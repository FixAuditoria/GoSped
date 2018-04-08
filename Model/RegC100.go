package Model

import "github.com/go-bongo/bongo"

// RegC100 : Documento - Nota Fiscal Eletrônica (código 55) e Nota Fiscal Eletrônica para Consumidor Final (código 65)
type RegC100 struct {
	bongo.DocumentBase `bson:",inline"`
	Reg                string
	IndOper            string
	IndEmit            string
	CodPart            string
	CodMod             string
	CodSit             string
	Ser                string
	NumDoc             string
	ChvNfe             string
	DtDoc              string
	DtES               string
	VlDoc              string
	IndPgto            string
	VlDesc             string
	VlAbatNt           string
	VlMerc             string
	IndFrt             string
	VlFrt              string
	VlSeg              string
	VlOutDa            string
	VlBcIcms           string
	VlIcms             string
	VlBcIcmsSt         string
	VlIcmsSt           string
	VlIpi              string
	VlPis              string
	VlCofins           string
	VlPisSt            string
	VlCofinsSt         string
	DtIni              string
	DtFin              string
	CnpjSped           string
}

// Populate : O métdodo é responsável por preencher os dados pelo sped
func (r *RegC100) Populate(l []string, reg0000 Reg0000) {
	r.Reg = l[1]
	r.IndOper = l[2]
	r.IndEmit = l[3]
	r.CodPart = l[4]
	r.CodMod = l[5]
	r.CodSit = l[6]
	r.Ser = l[7]
	r.NumDoc = l[8]
	r.ChvNfe = l[9]
	r.DtDoc = l[10]
	r.DtES = l[11]
	r.VlDoc = l[12]
	r.IndPgto = l[13]
	r.VlDesc = l[14]
	r.VlAbatNt = l[15]
	r.VlMerc = l[16]
	r.IndFrt = l[17]
	r.VlFrt = l[18]
	r.VlSeg = l[19]
	r.VlOutDa = l[20]
	r.VlBcIcms = l[21]
	r.VlIcms = l[22]
	r.VlBcIcmsSt = l[23]
	r.VlIcmsSt = l[24]
	r.VlIpi = l[25]
	r.VlPis = l[26]
	r.VlCofins = l[27]
	r.VlPisSt = l[28]
	r.VlCofinsSt = l[29]
	r.DtIni = reg0000.DtIni
	r.DtFin = reg0000.DtFin
	r.CnpjSped = reg0000.Cnpj
}
