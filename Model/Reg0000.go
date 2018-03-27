package Model

import (
	"gopkg.in/mgo.v2/bson"
)

// Reg0000 : Abertura do Arquivo Digital e Identificação da entidade
type Reg0000 struct {
	ID        bson.ObjectId `bson:"_id,omitempty" json:"id,omitempty"`
	Reg       string        `bson:"reg" json:"reg"`
	CodVer    string        `bson:"codver" json:"codver"`
	CodFin    string        `bson:"codfin" json:"codfin"`
	DtIni     string        `bson:"dtini" json:"dtini"`
	DtFin     string        `bson:"dtfin" json:"dtfin"`
	Nome      string        `bson:"nome" json:"nome"`
	Cnpj      string        `bson:"cnpj" json:"cnpj"`
	Cpf       string        `bson:"cpf" json:"cpf"`
	Uf        string        `bson:"uf" json:"uf"`
	Ie        string        `bson:"ie" json:"ie"`
	CodMun    string        `bson:"codmun" json:"codmun"`
	Im        string        `bson:"im" json:"im"`
	Suframa   string        `bson:"suframa" json:"suframa"`
	IndPerfil string        `bson:"indperfil" json:"indperfil"`
	IndAtiv   string        `bson:"indativ" json:"indativ"`
}

// Populate : O metodo é responsavel por preencher os dados pelo sped
func (r *Reg0000) Populate(l []string) {
	r.Reg = l[1]
	r.CodVer = l[2]
	r.CodFin = l[3]
	r.DtIni = l[4]
	r.DtFin = l[5]
	r.Nome = l[6]
	r.Cnpj = l[7]
	r.Cpf = l[8]
	r.Uf = l[9]
	r.Ie = l[10]
	r.CodMun = l[11]
	r.Im = l[12]
	r.Suframa = l[13]
	r.IndPerfil = l[14]
	r.IndAtiv = l[15]
}
