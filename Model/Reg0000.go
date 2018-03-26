package Model

import "time"

type Reg0000 struct {
	Reg       string
	CodVer    string
	CodFin    int
	DtIni     time.Time
	DtFin     time.Time
	Nome      string
	Cnpj      string
	Cpf       string
	Uf        string
	Ie        string
	CodMun    string
	Im        string
	Suframa   string
	IndPerfil string
	IndAtiv   int
}
