package Model

import (
	"gopkg.in/mgo.v2/bson"
)

// Reg0200 : Tabela de Identificação do Item (Produtos e Serviços)
type Reg0200 struct {
	ID         bson.ObjectId `bson:"_id,omitempty" json:"id,omitempty"`
	Reg        string        `bson:"reg" json:"reg"`
	CodItem    string        `bson:"coditem" json:"coditem"`
	DescrItem  string        `bson:"descritem" json:"descritem"`
	CodBarra   string        `bson:"codbarra" json:"codbarra"`
	CodAntItem string        `bson:"codantitem" json:"codantitem"`
	UnidInv    string        `bson:"unidinv" json:"unidinv"`
	TipoItem   string        `bson:"tipoitem" json:"tipoitem"`
	CodNcm     string        `bson:"codncm" json:"codncm"`
	ExIpi      string        `bson:"exipi" json:"exipi"`
	CodGen     string        `bson:"codgen" json:"codgen"`
	CodLst     string        `bson:"codlst" json:"codlst"`
	AliqIcms   string        `bson:"aliqicms" json:"aliqicms"`
	DtIni      string        `bson:"dtini" json:"dtini"`
	DtFin      string        `bson:"dtfin" json:"dtfin"`
	CnpjSped   string        `bson:"cnpj" json:"cnpj"`
}

// Populate: O métdodo é responsável por preencher os dados pelo sped
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
