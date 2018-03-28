package Controller

import (
	"fmt"
	"log"
	"path/filepath"
	"strings"

	"github.com/chapzin/GoSped/ConfigTom"
	"github.com/chapzin/GoSped/Dao"
	"github.com/chapzin/GoSped/Model"
	"github.com/chapzin/GoSped/Utilidades"
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
		// TODO : faz verificacao se é um arquivo sped fiscal e importa
		conteudo, err := Utilidades.ReadFile(arquivo)
		if err != nil {
			log.Fatalf("Erro:", err)
		}
		var reg0000 Model.Reg0000
		var reg0200 Model.Reg0200
		var regC100 Model.RegC100
		var controllerC100 RegC100Controller
		for _, linha := range conteudo {
			if len(linha) > 0 {
				if linha[:1] == "|" {
					l := strings.Split(linha, "|")
					if l[1] == "0000" {
						reg0000.Populate(l)
						err := reg0000Dao.Insert(reg0000)
						if err != nil {
							log.Panic(err)
						}
					}
					if l[1] == "0200" {
						reg0200.Reg = l[1]
						reg0200.CodItem = l[2]
						if len(reg0200.CodItem) != 10 {
							// codigo = reg0200.CodItem
							// fmt.Println(reg0200.CodItem + reg0000)
						}
					}
					// if codigo != "" {
					if l[1] == "C100" {
						regC100.Populate(l, reg0000)
						semchave := controllerC100.VerificarChave(regC100)
						if semchave != "" {
							fmt.Println(reg0000)
							fmt.Println(regC100)
							fmt.Println(semchave)
						}
						verifica2 := controllerC100.VerificarDtEmissaoChave(regC100)
						if verifica2 != "" {
							fmt.Println(verifica2)
						}

					}

					if l[1] == "C170" {
						if len(l[3]) != 10 {
							// fmt.Println(l[3] + ";" + chave)
						}

					}
					// }
				}
			}
		}
	}
	if extensao == ".csv" || extensao == ".CSV" {
		// TODO : faz verificacao se é um arquivo tipo siget e importa
	}
	if extensao == ".xml" || extensao == ".XML" {
		// TODO : faz verificacao se é um arquivo xml nfe, cte ou evento e importa
	}
}
