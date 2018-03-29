package Model

import (
	"gopkg.in/mgo.v2/bson"
)

// RegC100 : Documento - Nota Fiscal Eletrônica (código 55) e Nota Fiscal Eletrônica para Consumidor Final (código 65)
type RegC100 struct {
	ID         bson.ObjectId `bson:"_id,omitempty" json:"id,omitempty"`
	Reg        string        `bson:"reg" json:"reg"`
	IndOper    string        `bson:"indoper" json:"indoper"`
	IndEmit    string        `bson:"indemit" json:"indemit"`
	CodPart    string        `bson:"codpart" json:"codpart"`
	CodMod     string        `bson:"codmod" json:"codmod"`
	CodSit     string        `bson:"codsit" json:"codsit"`
	Ser        string        `bson:"ser" json:"ser"`
	NumDoc     string        `bson:"numdoc" json:"numdoc"`
	ChvNfe     string        `bson:"chvnfe" json:"chvnfe"`
	DtDoc      string        `bson:"dtdoc" json:"dtdoc"`
	DtES       string        `bson:"dtes" json:"dtes"`
	VlDoc      string        `bson:"vldoc" json:"vldoc"`
	IndPgto    string        `bson:"indpgto" json:"indpgto"`
	VlDesc     string        `bson:"vldesc" json:"vldesc"`
	VlAbatNt   string        `bson:"vlabatnt" json:"vlabatnt"`
	VlMerc     string        `bson:"vlmerc" json:"vlmerc"`
	IndFrt     string        `bson:"indfrt" json:"indfrt"`
	VlFrt      string        `bson:"vlfrt" json:"vlfrt"`
	VlSeg      string        `bson:"vlseg" json:"vlseg"`
	VlOutDa    string        `bson:"vloutda" json:"vloutda"`
	VlBcIcms   string        `bson:"vlbcicms" json:"vlbcicms"`
	VlIcms     string        `bson:"vlicms" json:"vlicms"`
	VlBcIcmsSt string        `bson:"vlbcicmsst" json:"vlbcicmsst"`
	VlIcmsSt   string        `bson:"vlicmsst" json:"vlicmsst"`
	VlIpi      string        `bson:"vlipi" json:"vlipi"`
	VlPis      string        `bson:"vlpis" json:"vlpis"`
	VlCofins   string        `bson:"vlcofins" json:"vlcofins"`
	VlPisSt    string        `bson:"vlpisst" json:"vlpisst"`
	VlCofinsSt string        `bson:"vlcofinsst" json:"vlcofinsst"`
	DtIni      string        `bson:"dtini" json:"dtini"`
	DtFin      string        `bson:"dtfin" json:"dtfin"`
	CnpjSped   string        `bson:"cnpjsped" json:"cnpjsped"`
}

// Populate: O métdodo é responsável por preencher os dados pelo sped
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
