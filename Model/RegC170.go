package Model

import "github.com/go-bongo/bongo"

// RegC170 : Complemento de Documento - Itens do Documento (código 01, 1B, 04 e 55)
type RegC170 struct {
	bongo.DocumentBase `bson:",inline"`
	Reg                string
	NumItem            string
	CodItem            string
	DescrCompl         string
	Qtd                string
	Unid               string
	VlItem             string
	VlDesc             string
	IndMov             string
	CstIcms            string
	Cfop               string
	CodNat             string
	VlBcIcms           string
	AliqIcms           string
	VlIcms             string
	VlBcIcmsSt         string
	AliqSt             string
	VlIcmsSt           string
	IndApur            string
	CstIpi             string
	CodEnq             string
	VlBcIpi            string
	AliqIpi            string
	VlIpi              string
	CstPis             string
	VlBcPis            string
	AliqPis01          string
	QuantBcPis         string
	AliqPis02          string
	VlPis              string
	CstCofins          string
	VlBcCofins         string
	AliqCofins01       string
	QuantBcCofins      string
	AliqCofins02       string
	VlCofins           string
	CodCta             string
	EntradaSaida       string
	NumDoc             string
	ChvNfe             string
	DtIni              string
	DtFin              string
	CnpjSped           string
}

// Populate : O métdodo é responsável por preencher os dados pelo sped
func (r *RegC170) Populate(l []string, reg0000 Reg0000, regC100 RegC100) {
	r.Reg = l[1]
	r.NumItem = l[2]
	r.CodItem = l[3]
	r.DescrCompl = l[4]
	r.Qtd = l[5]
	r.Unid = l[6]
	r.VlItem = l[7]
	r.VlDesc = l[8]
	r.IndMov = l[9]
	r.CstIcms = l[10]
	r.Cfop = l[11]
	r.CodNat = l[12]
	r.VlBcIcms = l[13]
	r.AliqIcms = l[14]
	r.VlIcms = l[15]
	r.VlBcIcmsSt = l[16]
	r.AliqSt = l[17]
	r.VlIcmsSt = l[18]
	r.IndApur = l[19]
	r.CstIpi = l[20]
	r.CodEnq = l[21]
	r.VlBcIpi = l[22]
	r.AliqIpi = l[23]
	r.VlIpi = l[24]
	r.CstPis = l[25]
	r.VlBcPis = l[26]
	r.AliqPis01 = l[27]
	r.QuantBcPis = l[28]
	r.AliqPis02 = l[29]
	r.VlPis = l[30]
	r.CstCofins = l[31]
	r.VlBcCofins = l[32]
	r.AliqCofins01 = l[33]
	r.QuantBcCofins = l[34]
	r.AliqCofins02 = l[35]
	r.VlCofins = l[36]
	r.CodCta = l[37]
	r.EntradaSaida = regC100.DtES
	r.NumDoc = regC100.NumDoc
	r.ChvNfe = regC100.ChvNfe
	r.DtIni = reg0000.DtIni
	r.DtFin = reg0000.DtFin
	r.CnpjSped = reg0000.Cnpj
}
