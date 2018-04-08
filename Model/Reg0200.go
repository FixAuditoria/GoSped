package Model

import "github.com/go-bongo/bongo"

// Reg0200 : Tabela de Identificação do Item (Produtos e Serviços)
type Reg0200 struct {
	bongo.DocumentBase `bson:",inline"`
	Reg                string
	CodItem            string
	DescrItem          string
	CodBarra           string
	CodAntItem         string
	UnidInv            string
	TipoItem           string
	CodNcm             string
	ExIpi              string
	CodGen             string
	CodLst             string
	AliqIcms           string
	DtIni              string
	DtFin              string
	CnpjSped           string
}

// Populate : O métdodo é responsável por preencher os dados pelo sped
func (r *Reg0200) Populate(l []string, reg0000 Reg0000) {
	r.Reg = l[1]
	r.CodItem = l[2]
	r.DescrItem = l[3]
	r.CodBarra = l[4]
	r.CodAntItem = l[5]
	r.UnidInv = l[6]
	r.TipoItem = l[7]
	r.CodNcm = l[8]
	r.ExIpi = l[9]
	r.CodGen = l[10]
	r.CodLst = l[11]
	r.AliqIcms = l[12]
	r.DtIni = reg0000.DtIni
	r.DtFin = reg0000.DtFin
	r.CnpjSped = reg0000.Cnpj
}
