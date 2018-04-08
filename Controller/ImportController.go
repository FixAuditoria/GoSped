package Controller

import (
	"path/filepath"

	"github.com/go-bongo/bongo"

	"github.com/chapzin/GoSped/Utilidades"
)

// ImportController : Responsavel por separar e executar os comandos nos arquivos txt, csv e xml
type ImportController struct {
	arquivo string
}

// Importar : metodo responsavel pela separacao dos arquivos
func (i *ImportController) Importar(path string, conn *bongo.Connection) {
	arquivos, _ := Utilidades.ListFiles(path)
	for _, arq := range arquivos {
		extensao := filepath.Ext(arq)
		if extensao == ".txt" || extensao == ".TXT" {
			var spedcontroller SpedController
			spedcontroller.addDB(arq, conn)
			// spedcontroller.validacoes(arq)
		}
		if extensao == ".csv" || extensao == ".CSV" {
			// TODO : faz verificacao se é um arquivo tipo siget e importa
		}
		if extensao == ".xml" || extensao == ".XML" {
			// TODO : faz verificacao se é um arquivo xml nfe, cte ou evento e importa
			// fmt.Println(arq)
			ListaXmls(arq, conn)
		}
	}
}
