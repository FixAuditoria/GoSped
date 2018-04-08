package Model

import "github.com/go-bongo/bongo"

// Reg0100 : Dados do Contabilista
type Reg0100 struct {
	bongo.DocumentBase `bson:",inline"`
	Reg                string
	Nome               string
	Cpf                string
	Crc                string
	Cnpj               string
	Cep                string
	End                string
	Num                string
	Compl              string
	Bairro             string
	Fone               string
	Fax                string
	Email              string
	CodMun             string
	DtIni              string
	DtFin              string
	CnpjSped           string
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
