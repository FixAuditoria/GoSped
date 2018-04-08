package Model

import (
	"github.com/jinzhu/gorm"
)

// Reg0220 : Fatores de Conversão de Unidades
type Reg0220 struct {
	gorm.Model
	Reg      string `gorm:"type:varchar(4)" bson:"reg" json:"reg"`
	UnidConv string `gorm:"type:varchar(6)" bson:"unidconv" json:"unidconv"`
	FatConv  string `gorm:"type:varchar(20)" bson:"fatconv" json:"fatconv"`
	UnidCod  string `gorm:"type:varchar(10)" bson:"unidcod" json:"unidcod"`
	CodItem  string `gorm:"type:varchar(60)" bson:"coditem" json:"coditem"`
	DtIni    string `gorm:"type:varchar(8)" bson:"dtini" json:"dtini"`
	DtFin    string `gorm:"type:varchar(8)" bson:"dtfin" json:"dtfin"`
	CnpjSped string `gorm:"type:varchar(14)" bson:"cnpjsped" json:"cnpjsped"`
}

// TableName : Funcao responsavel por definir o nome na tabela
func (Reg0220) TableName() string {
	return "reg0220"
}

// CreateDB : funcao para criar o banco de dados do Registro 0100
func (r *Reg0220) CreateDB(db gorm.DB) {
	db.AutoMigrate(r)
}

// DropDB : funcao para apagar banco de dados do Registro 0100
func (r *Reg0220) DropDB(db gorm.DB) {
	db.DropTable(r)
}

// Populate : O métdodo é responsável por preencher os dados pelo sped
func (r *Reg0220) Populate(l []string, reg0000 Reg0000, reg0200 Reg0200) {
	r.Reg = l[1]
	r.UnidConv = l[2]
	r.FatConv = l[3]
	r.UnidCod = reg0200.UnidInv
	r.CodItem = reg0200.CodItem
	r.DtIni = reg0000.DtIni
	r.DtFin = reg0000.DtFin
	r.CnpjSped = reg0000.Cnpj
}
