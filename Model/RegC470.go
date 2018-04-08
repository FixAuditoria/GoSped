package Model

import "github.com/go-bongo/bongo"

// RegC470 : Itens do Documento Fiscal Emitido por ECF (código 02 e 2D)
type RegC470 struct {
	bongo.DocumentBase `bson:",inline"`
	Reg                string
	CodItem            string
	Qtd                string
	QtdCanc            string
	Unid               string
	VlItem             string
	CstIcms            string
	Cfop               string
	AliqIcms           string
	VlPis              string
	VlCofins           string
	DtIni              string
	DtFin              string
	CnpjSped           string
}

// Populate : O métdodo é responsável por preencher os dados pelo sped
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
