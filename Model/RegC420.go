package Model

import (
	"gopkg.in/mgo.v2/bson"
)

// RegC420 : Registro dos Totalizadores Parciais da Redução Z (código 02, 2D e 60)
type RegC420 struct {
	ID         bson.ObjectId `bson:"_id,omitempty" json:"id,omitempty"`
	Reg        string        `bson:"reg" json:"reg"`
	CodTotPar  string        `bson:"codtotpar" json:"codtotpar"`
	VlrAcumTot string        `bson:"vlracumtot" json:"vlracumtot"`
	NrTot      string        `bson:"nrtot" json:"nrtot"`
	DescrNrTot string        `bson:"descrnrtot" json:"descrnrtot"`
	DtIni      string        `bson:"dtini" json:"dtini"`
	DtFin      string        `bson:"dtfin" json:"dtfin"`
	CnpjSped   string        `bson:"cnpjsped" json:"cnpjsped"`
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
