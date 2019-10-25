package Controller

import (
	"path/filepath"

	"github.com/go-bongo/bongo"

	"fmt"
	"os"

	"github.com/FixAuditoria/GoSped/Utilidades"
	"github.com/mholt/archiver"
)

// ImportController : Responsavel por separar e executar os comandos nos arquivos txt, csv e xml
type ImportController struct {
	arquivo string
}

// Importar : metodo responsavel pela separacao dos arquivos
func (i *ImportController) Importar(path string, conn *bongo.Connection) {
	arquivos, _ := Utilidades.ListarArquivosV2(path)
	for _, arq := range arquivos {
		extensao := filepath.Ext(arq)
		DeleteArquivoVazio(arq)
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
			ProcessarXmls(arq, conn)
		}
		if extensao == ".pdf" || extensao == ".PDF" {
			os.Remove(arq)
		}

		if extensao == ".rar" || extensao == ".RAR" {
			err := archiver.Unarchive(arq, "importar/")
			if err != nil {
				fmt.Println(err.Error())
			} else {
				os.Remove(arq)
			}
		}

		if extensao == ".zip" || extensao == ".ZIP" {
			err := archiver.Unarchive(arq, "importar/")
			if err != nil {
				fmt.Println(err.Error())
			} else {
				os.Remove(arq)
			}
		}

		if extensao == ".exe" || extensao == ".EXE" {
			os.Remove(arq)
		}

		if extensao == ".html" || extensao == ".HTML" || extensao == ".htm" || extensao == ".HTM" {
			os.Remove(arq)
		}

		if extensao == ".err" || extensao == ".ERR" {
			os.Remove(arq)
		}

	}
}
