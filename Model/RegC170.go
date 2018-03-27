package Model

import (
	"gopkg.in/mgo.v2/bson"
)

// RegC170 : Complemento de Documento - Itens do Documento (código 01, 1B, 04 e 55)
type RegC170 struct {
	ID            bson.ObjectId `bson:"_id,omitempty" json:"id,omitempty"`
	Reg           string        `bson:"reg" json:"reg"`
	NumItem       string        `bson:"numitem" json:"numitem"`
	CodItem       string        `bson:"coditem" json:"coditem"`
	DescrCompl    string        `bson:"descrcompl" json:"descrcompl"`
	Qtd           string        `bson:"qtd" json:"qtd"`
	Unid          string        `bson:"unid" json:"unid"`
	VlItem        string        `bson:"vlitem" json:"vlitem"`
	VlDesc        string        `bson:"vldesc" json:"vldesc"`
	IndMov        string        `bson:"indmov" json:"indmov"`
	CstIcms       string        `bson:"csticms" json:"csticms"`
	Cfop          string        `bson:"cfop" json:"cfop"`
	CodNat        string        `bson:"codnat" json:"codnat"`
	VlBcIcms      string        `bson:"vlbcicms" json:"vlbcicms"`
	AliqIcms      string        `bson:"aliqicms" json:"aliqicms"`
	VlIcms        string        `bson:"vlicms" json:"vlicms"`
	VlBcIcmsSt    string        `bson:"vlbicmsst" json:"vlbicmsst"`
	AliqSt        string        `bson:"aliqst" json:"aliqst"`
	VlIcmsSt      string        `bson:"vlicmsst" json:"vlicmsst"`
	IndApur       string        `bson:"indapur" json:"indapur"`
	CstIpi        string        `bson:"cstipi" json:"cstipi"`
	CodEnq        string        `bson:"codenq" json:"codenq"`
	VlBcIpi       string        `bson:"vlbcipi" json:"vlbcipi"`
	AliqIpi       string        `bson:"aliqipi" json:"aliqipi"`
	VlIpi         string        `bson:"vlipi" json:"vlipi"`
	CstPis        string        `bson:"cstpis" json:"cstpis"`
	VlBcPis       string        `bson:"vlbcpis" json:"vlbcpis"`
	AliqPis01     string        `bson:"aliqpis01" json:"aliqpis01"`
	QuantBcPis    string        `bson:"quantbcpis" json:"quantbcpis"`
	AliqPis02     string        `bson:"aliqpis02" json:"aliqpis02"`
	VlPis         string        `bson:"vlpis" json:"vlpis"`
	CstCofins     string        `bson:"cstcofins" json:"cstcofins"`
	VlBcCofins    string        `bson:"vlbccofins" json:"vlbccofins"`
	AliqCofins01  string        `bson:"aliqcofins01" json:"aliqcofins01"`
	QuantBcCofins string        `bson:"quantbccofins" json:"quantbccofins"`
	AliqCofins02  string        `bson:"aliqcofins02" json:"aliqcofins02"`
	VlCofins      string        `bson:"vlcofins" json:"vlcofins"`
	CodCta        string        `bson:"codcta" json:"codcta"`
	EntradaSaida  string        `bson:"entradasaida" json:"entradasaida"`
	NumDoc        string        `bson:"numdoc" json:"numdoc"`
	DtIni         string        `bson:"dtini" json:"dtini"`
	DtFin         string        `bson:"dtfin" json:"dtfin"`
	Cnpj          string        `bson:"cnpj" json:"cnpj"`
}

// Populate: O métdodo é responsável por preencher os dados pelo sped
func (r *RegC170) Populate(l []string, reg0000 Reg0000) {
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
	r.EntradaSaida = l[38]
	r.NumDoc = l[39]
	r.DtIni = reg0000.DtIni
	r.DtFin = reg0000.DtFin
	r.Cnpj = reg0000.Cnpj
}
