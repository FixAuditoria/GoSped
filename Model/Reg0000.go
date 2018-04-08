package Model

import "github.com/go-bongo/bongo"

// Reg0000 : Abertura do Arquivo Digital e Identificação da entidade
type Reg0000 struct {
	bongo.DocumentBase `bson:",inline"`
	Reg                string
	CodVer             string
	CodFin             string
	DtIni              string
	DtFin              string
	Nome               string
	Cnpj               string
	Cpf                string
	Uf                 string
	Ie                 string
	CodMun             string
	Im                 string
	Suframa            string
	IndPerfil          string
	IndAtiv            string
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
