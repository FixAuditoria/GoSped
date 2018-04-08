package Model

import "github.com/go-bongo/bongo"

// Reg0220 : Fatores de Conversão de Unidades
type Reg0220 struct {
	bongo.DocumentBase `bson:",inline"`
	Reg                string
	UnidConv           string
	FatConv            string
	UnidCod            string
	CodItem            string
	DtIni              string
	DtFin              string
	CnpjSped           string
}

// Populate : O métdodo é responsável por preencher os dados pelo sped
func (r *Reg0220) Populate(l []string, reg0000 Reg0000, reg0200 Reg0200) {
	r.Reg = l[1]
	r.UnidConv = l[2]
	r.FatConv = l[3]
	r.UnidCod = reg0200.UnidInv
	r.CodItem = reg0200.CodItem
	r.DtIni = reg0000.DtIni
	r.DtFin = reg0000.DtFin
	r.CnpjSped = reg0000.Cnpj
}
