package Model

import (
	"gopkg.in/mgo.v2/bson"
)

// Reg0150 : Tabela de Cadastro do Participante
type Reg0150 struct {
	ID       bson.ObjectId `bson:"_id,omitempty" json:"id,omitempty"`
	Reg      string        `bson:"reg" json:"reg"`
	CodPart  string        `bson:"codpart" json:"codpart"`
	Nome     string        `bson:"nome" json:"nome"`
	CodPais  string        `bson:"codpais" json:"codpais"`
	Cnpj     string        `bson:"cnpj" json:"cnpj"`
	Cpf      string        `bson:"cpf" json:"cpf"`
	Ie       string        `bson:"ie" json:"ie"`
	CodMun   string        `bson:"codmun" json:"codmun"`
	Suframa  string        `bson:"suframa" json:"suframa"`
	Endereco string        `bson:"endereco" json:"endereco"`
	Num      string        `bson:"num" json:"num"`
	Compl    string        `bson:"compl" json:"compl"`
	Bairro   string        `bson:"bairro" json:"bairro"`
	DtIni    string        `bson:"dtini" json:"dtini"`
	DtFin    string        `bson:"dtfin" json:"dtfin"`
	CnpjSped string        `bson:"cnpjsped" json:"cnpjsped"`
}

// Populate: O métdodo é responsável por preencher os dados pelo sped
func (r *Reg0150) Populate(l []string, reg0000 Reg0000) {
	r.Reg = l[1]
	r.CodPart = l[2]
	r.Nome = l[3]
	r.CodPais = l[4]
	r.Cnpj = l[5]
	r.Cpf = l[6]
	r.Ie = l[7]
	r.CodMun = l[8]
	r.Suframa = l[9]
	r.Endereco = l[10]
	r.Num = l[11]
	r.Compl = l[12]
	r.Bairro = l[13]
	r.DtIni = reg0000.DtIni
	r.DtFin = reg0000.DtFin
	r.CnpjSped = reg0000.Cnpj
}
