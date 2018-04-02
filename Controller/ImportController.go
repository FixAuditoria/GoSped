package Controller

import (
	"path/filepath"
)

// ImportController : Responsavel por separar e executar os comandos nos arquivos txt, csv e xml
type ImportController struct {
	arquivo string
}

// Importar : metodo responsavel pela separacao dos arquivos
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
