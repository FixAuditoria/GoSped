package Model

import "github.com/go-bongo/bongo"

// Reg0190 : Identificação das unidades de medida
type Reg0190 struct {
	bongo.DocumentBase `bson:",inline"`
	Reg                string
	Unid               string
	Descr              string
	DtIni              string
	DtFin              string
	CnpjSped           string
}

// Populate : O métdodo é responsável por preencher os dados pelo sped
func (r *Reg0190) Populate(l []string, reg0000 Reg0000) {
	r.Reg = l[1]
	r.Unid = l[2]
	r.Descr = l[3]
	r.DtIni = reg0000.DtIni
	r.DtFin = reg0000.DtFin
	r.CnpjSped = reg0000.Cnpj
}
