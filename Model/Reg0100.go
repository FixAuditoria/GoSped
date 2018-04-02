package Model

import (
	"gopkg.in/mgo.v2/bson"
)

// Reg0100 : Dados do Contabilista
type Reg0100 struct {
	ID       bson.ObjectId `bson:"_id,omitempty" json:"id,omitempty"`
	Reg      string        `bson:"reg" json:"reg"`
	Nome     string        `bson:"nome" json:"nome"`
	Cpf      string        `bson:"cpf" json:"cpf"`
	Crc      string        `bson:"crc" json:"crc"`
	Cnpj     string        `bson:"cnpj" json:"cnpj"`
	Cep      string        `bson:"cep" json:"cep"`
	End      string        `bson:"end" json:"end"`
	Num      string        `bson:"num" json:"num"`
	Compl    string        `bson:"compl" json:"compl"`
	Bairro   string        `bson:"bairro" json:"bairro"`
	Fone     string        `bson:"fone" json:"fone"`
	Fax      string        `bson:"fax" json:"fax"`
	Email    string        `bson:"email" json:"email"`
	CodMun   string        `bson:"codnum" json:"codnum"`
	DtIni    string        `bson:"dtini" json:"dtini"`
	DtFin    string        `bson:"dtfin" json:"dtfin"`
	CnpjSped string        `bson:"cnpjsped" json:"cnpjsped"`
}

//Populate : O métdodo é responsável por preencher os dados pelo sped
func (r *Reg0100) Populate(l []string, reg0000 Reg0000) {
	r.Reg = l[1]
	r.Nome = l[2]
	r.Cpf = l[3]
	r.Crc = l[4]
	r.Cnpj = l[5]
	r.Cep = l[6]
	r.End = l[7]
	r.Num = l[8]
	r.Compl = l[9]
	r.Bairro = l[10]
	r.Fone = l[11]
	r.Fax = l[12]
	r.Email = l[13]
	r.CodMun = l[14]
	r.DtIni = reg0000.DtIni
	r.DtFin = reg0000.DtFin
	r.CnpjSped = reg0000.Cnpj
}
