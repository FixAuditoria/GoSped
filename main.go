package main

import (
	"fmt"
	"log"
	"strings"

	. "github.com/chapzin/GoSped/ConfigTom"
	. "github.com/chapzin/GoSped/Controller"
	. "github.com/chapzin/GoSped/Dao"
	. "github.com/chapzin/GoSped/Model"
	. "github.com/chapzin/GoSped/Utilidades"
)

var cofing2 = ConfigSped{}
var reg0000Dao = Reg0000Dao{}

func init() {
	cofing2.Read()
	reg0000Dao.Server = cofing2.Server
	reg0000Dao.Database = cofing2.Database
	reg0000Dao.Connect()

}

func main() {

	arquivos, _ := ListFiles("./speds/")
	for _, arq := range arquivos {
		conteudo, err := ReadFile(arq)
		if err != nil {
			log.Fatalf("Erro:", err)
		}
		var reg0000 Reg0000
		var reg0200 Reg0200
		var regC100 RegC100
		var controllerC100 RegC100Controller
		var importar ImportController
		importar.Importar(arq)
		// var codigo string
		// var chave string
		for _, linha := range conteudo {
			// fmt.Println(ind, linha)

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

}
