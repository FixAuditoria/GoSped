package Model

import "github.com/go-bongo/bongo"

// RegC405 : Redução Z (código 02, 2D e 60)
type RegC405 struct {
	bongo.DocumentBase `bson:",inline"`
	Reg                string
	DtDoc              string
	Cro                string
	Crz                string
	NumCooFin          string
	GtFin              string
	VlBrt              string
	DtIni              string
	DtFin              string
	CnpjSped           string
}

// Populate : O métdodo é responsável por preencher os dados pelo sped
func (r *RegC405) Populate(l []string, reg0000 Reg0000) {
	r.Reg = l[1]
	r.DtDoc = l[2]
	r.Cro = l[3]
	r.Crz = l[4]
	r.NumCooFin = l[5]
	r.GtFin = l[6]
	r.VlBrt = l[7]
	r.DtIni = reg0000.DtIni
	r.DtFin = reg0000.DtFin
	r.CnpjSped = reg0000.Cnpj
}
