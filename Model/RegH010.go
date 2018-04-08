package Model

import "github.com/go-bongo/bongo"

// RegH010 : Inventário
type RegH010 struct {
	bongo.DocumentBase `bson:",inline"`
	Reg                string
	CodItem            string
	Unid               string
	Qtd                string
	VlUnit             string
	VlItem             string
	IndProp            string
	CodPart            string
	TxtCompl           string
	CodCta             string
	VlItemIr           string
	DtInv              string
	DtIni              string
	DtFin              string
	CnpjSped           string
}

// Populate : O métdodo é responsável por preencher os dados pelo sped
func (r *RegH010) Populate(l []string, reg0000 Reg0000, regH005 RegH005) {
	r.Reg = l[1]
	r.CodItem = l[2]
	r.Unid = l[3]
	r.Qtd = l[4]
	r.VlUnit = l[5]
	r.VlItem = l[6]
	r.IndProp = l[7]
	r.CodPart = l[8]
	r.TxtCompl = l[9]
	r.CodCta = l[10]
	r.VlItemIr = l[11]
	r.DtInv = regH005.DtInv
	r.DtIni = reg0000.DtIni
	r.DtFin = reg0000.DtFin
	r.CnpjSped = reg0000.Cnpj
}
