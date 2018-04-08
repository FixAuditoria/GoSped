package Model

import "github.com/go-bongo/bongo"

// RegC400 : Equipamento ECF (código 02, 2D e 60),
type RegC400 struct {
	bongo.DocumentBase `bson:",inline"`
	Reg                string
	CodMod             string
	EcfMod             string
	EcfFab             string
	EcfCx              string
	DtIni              string
	DtFin              string
	CnpjSped           string
}

// Populate : O métdodo é responsável por preencher os dados pelo sped
func (r *RegC400) Populate(l []string, reg0000 Reg0000) {
	r.Reg = l[1]
	r.CodMod = l[2]
	r.EcfMod = l[3]
	r.EcfFab = l[4]
	r.EcfCx = l[5]
	r.DtIni = reg0000.DtIni
	r.DtFin = reg0000.DtFin
	r.CnpjSped = reg0000.Cnpj
}
