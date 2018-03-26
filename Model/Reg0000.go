package Model

import "time"

// Reg0000 : Abertura do Arquivo Digital e Identificação da entidade
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
