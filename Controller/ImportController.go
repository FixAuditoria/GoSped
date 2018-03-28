package Controller

import (
	"path/filepath"

	"github.com/chapzin/GoSped/ConfigTom"
	"github.com/chapzin/GoSped/Dao"
)

var cofing2 = ConfigTom.ConfigSped{}
var reg0000Dao = Dao.Reg0000Dao{}

func init() {
	cofing2.Read()
	reg0000Dao.Server = cofing2.Server
	reg0000Dao.Database = cofing2.Database
	reg0000Dao.Connect()

}

type ImportController struct {
	arquivo string
}

func (i *ImportController) Importar(arquivo string) {
	extensao := filepath.Ext(arquivo)
	if extensao == ".txt" || extensao == ".TXT" {
		var spedcontroller SpedController
		spedcontroller.addMongo(arquivo)
		spedcontroller.validacoes(arquivo)
	}
	if extensao == ".csv" || extensao == ".CSV" {
		// TODO : faz verificacao se é um arquivo tipo siget e importa
	}
	if extensao == ".xml" || extensao == ".XML" {
		// TODO : faz verificacao se é um arquivo xml nfe, cte ou evento e importa
	}
}
