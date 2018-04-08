package Model

import "github.com/go-bongo/bongo"

// RegC420 : Registro dos Totalizadores Parciais da Redução Z (código 02, 2D e 60)
type RegC420 struct {
	bongo.DocumentBase `bson:",inline"`
	Reg                string
	CodTotPar          string
	VlrAcumTot         string
	NrTot              string
	DescrNrTot         string
	DtIni              string
	DtFin              string
	CnpjSped           string
}

// Populate : O métdodo é responsável por preencher os dados pelo sped
func (r *RegC420) Populate(l []string, reg0000 Reg0000) {
	r.Reg = l[1]
	r.CodTotPar = l[2]
	r.VlrAcumTot = l[3]
	r.NrTot = l[4]
	r.DescrNrTot = l[5]
	r.DtIni = reg0000.DtIni
	r.DtFin = reg0000.DtFin
	r.CnpjSped = reg0000.Cnpj
}
