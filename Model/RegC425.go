package Model

import "github.com/go-bongo/bongo"

// RegC425 : Resumo de itens do movimento diário (código 02 e 2D)
type RegC425 struct {
	bongo.DocumentBase `bson:",inline"`
	Reg                string
	CodItem            string
	Qtd                string
	Unid               string
	VlItem             string
	VlPis              string
	VlCofins           string
	DtVenda            string
	DtIni              string
	DtFin              string
	CnpjSped           string
}

// Populate : O métdodo é responsável por preencher os dados pelo sped
func (r *RegC425) Populate(l []string, reg0000 Reg0000, regC405 RegC405) {
	r.Reg = l[1]
	r.CodItem = l[2]
	r.Qtd = l[3]
	r.Unid = l[4]
	r.VlItem = l[5]
	r.VlPis = l[6]
	r.VlCofins = l[7]
	r.DtVenda = regC405.DtDoc
	r.DtIni = reg0000.DtIni
	r.DtFin = reg0000.DtFin
	r.CnpjSped = reg0000.Cnpj
}
