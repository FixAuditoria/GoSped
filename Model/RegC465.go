package Model

import "github.com/go-bongo/bongo"

// RegC465 : Complemento do Cupom Fiscal Eletrônico Emitido por ECF - CF-e-ECF (código 60)
type RegC465 struct {
	bongo.DocumentBase `bson:",inline"`
	Reg                string
	ChvCfe             string
	NumCcF             string
	DtIni              string
	DtFin              string
	CnpjSped           string
}

// Populate : O métdodo é responsável por preencher os dados pelo sped
func (r *RegC465) Populate(l []string, reg0000 Reg0000) {
	r.Reg = l[1]
	r.ChvCfe = l[2]
	r.NumCcF = l[3]
	r.DtIni = reg0000.DtIni
	r.DtFin = reg0000.DtFin
	r.CnpjSped = reg0000.Cnpj
}
