package Model

import "github.com/go-bongo/bongo"

// RegC460 : Documento Fiscal Emitido por ECF (código 02, 2D e 60)
type RegC460 struct {
	bongo.DocumentBase `bson:",inline"`
	Reg                string
	CodMod             string
	CodSit             string
	NumDoc             string
	DtDoc              string
	VlDoc              string
	VlPis              string
	VlCofins           string
	CpfCnpj            string
	NomAdq             string
	DtIni              string
	DtFin              string
	CnpjSped           string
}

// Populate : O métdodo é responsável por preencher os dados pelo sped
func (r *RegC460) Populate(l []string, reg0000 Reg0000) {
	r.Reg = l[1]
	r.CodMod = l[2]
	r.CodSit = l[3]
	r.NumDoc = l[4]
	r.DtDoc = l[5]
	r.VlDoc = l[6]
	r.VlPis = l[7]
	r.VlCofins = l[8]
	r.CpfCnpj = l[9]
	r.NomAdq = l[10]
	r.DtIni = reg0000.DtIni
	r.DtFin = reg0000.DtFin
	r.CnpjSped = reg0000.Cnpj
}
