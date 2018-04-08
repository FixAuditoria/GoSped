package Model

import "github.com/go-bongo/bongo"

// Reg0150 : Tabela de Cadastro do Participante
type Reg0150 struct {
	bongo.DocumentBase `bson:",inline"`
	Reg                string
	CodPart            string
	Nome               string
	CodPais            string
	Cnpj               string
	Cpf                string
	Ie                 string
	CodMun             string
	Suframa            string
	Endereco           string
	Num                string
	Compl              string
	Bairro             string
	DtIni              string
	DtFin              string
	CnpjSped           string
}

// Populate : O métdodo é responsável por preencher os dados pelo sped
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
