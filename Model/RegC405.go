package Model

import (
	"gopkg.in/mgo.v2/bson"
)

// RegC405 : Redução Z (código 02, 2D e 60)
type RegC405 struct {
	ID        bson.ObjectId `bson:"_id,omitempty" json:"id,omitempty"`
	Reg       string        `bson:"reg" json:"reg"`
	DtDoc     string        `bson:"dtdoc" json:"dtdoc"`
	Cro       string        `bson:"cro" json:"cro"`
	Crz       string        `bson:"crz" json:"crz"`
	NumCooFin string        `bson:"numcoofin" json:"numcoofin"`
	GtFin     string        `bson:"gtfin" json:"gtfin"`
	VlBrt     string        `bson:"vlbrt" json:"vlbrt"`
	DtIni     string        `bson:"dtini" json:"dtini"`
	DtFin     string        `bson:"dtfin" json:"dtfin"`
	CnpjSped  string        `bson:"cnpjsped" json:"cnpjsped"`
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
