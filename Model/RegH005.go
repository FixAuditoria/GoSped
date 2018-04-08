package Model

import "github.com/go-bongo/bongo"

// RegH005 : Totais do Inventário
type RegH005 struct {
	bongo.DocumentBase `bson:",inline"`
	Reg                string
	DtInv              string
	VlInv              string
	MotInv             string
	DtIni              string
	DtFin              string
	CnpjSped           string
}

// Populate : O métdodo é responsável por preencher os dados pelo sped
func (r *RegH005) Populate(l []string, reg0000 Reg0000) {
	r.Reg = l[1]
	r.DtInv = l[2]
	r.VlInv = l[3]
	r.MotInv = l[4]
	r.DtIni = reg0000.DtIni
	r.DtFin = reg0000.DtFin
	r.CnpjSped = reg0000.Cnpj
}
