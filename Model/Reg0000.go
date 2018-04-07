package Model

import (
	"github.com/jinzhu/gorm"
)

// Reg0000 : Abertura do Arquivo Digital e Identificação da entidade
type Reg0000 struct {
	gorm.Model
	Reg       string `gorm:"type:varchar(4)" bson:"reg" json:"reg"`
	CodVer    string `gorm:"type:varchar(3)" bson:"codver" json:"codver"`
	CodFin    string `gorm:"type:varchar(3)" bson:"codfin" json:"codfin"`
	DtIni     string `gorm:"type:varchar(8)" bson:"dtini" json:"dtini"`
	DtFin     string `gorm:"type:varchar(8)" bson:"dtfin" json:"dtfin"`
	Nome      string `gorm:"type:varchar(100)" bson:"nome" json:"nome"`
	Cnpj      string `gorm:"type:varchar(14)" bson:"cnpj" json:"cnpj"`
	Cpf       string `gorm:"type:varchar(11)" bson:"cpf" json:"cpf"`
	Uf        string `gorm:"type:varchar(2)" bson:"uf" json:"uf"`
	Ie        string `gorm:"type:varchar(14)" bson:"ie" json:"ie"`
	CodMun    string `gorm:"type:varchar(7)" bson:"codmun" json:"codmun"`
	Im        string `gorm:"type:varchar(15)" bson:"im" json:"im"`
	Suframa   string `gorm:"type:varchar(9)" bson:"suframa" json:"suframa"`
	IndPerfil string `gorm:"type:varchar(1)" bson:"indperfil" json:"indperfil"`
	IndAtiv   string `gorm:"type:varchar(5)" bson:"indativ" json:"indativ"`
}

// TableName : definindo nome da tabela no Registro 0000
func (Reg0000) TableName() string {
	return "reg0000"
}

// CreateDB : funcao para criar o banco de dados do Registro 0000
func (r *Reg0000) CreateDB(db gorm.DB) {
	db.AutoMigrate(r)
}

// DropDB : funcao para apagar banco de dados do Registro 0000
func (r *Reg0000) DropDB(db gorm.DB) {
	db.DropTable(r)
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
