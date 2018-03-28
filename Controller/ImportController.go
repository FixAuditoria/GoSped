package Controller

import (
	"path/filepath"
)

type ImportController struct {
	arquivo string
}

func (i *ImportController) Importar(arquivo string) {
	extensao := filepath.Ext(arquivo)
	if extensao == ".txt" || extensao == ".TXT" {
		// TODO : faz verificacao se é um arquivo sped fiscal e importa
	}
	if extensao == ".csv" || extensao == ".CSV" {
		// TODO : faz verificacao se é um arquivo tipo siget e importa
	}
	if extensao == ".xml" || extensao == ".XML" {
		// TODO : faz verificacao se é um arquivo xml nfe, cte ou evento e importa
	}
}
