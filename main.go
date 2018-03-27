package main

import (
	"fmt"
	"log"
	"strings"

	. "github.com/chapzin/GoSped/ConfigTom"
	. "github.com/chapzin/GoSped/Dao"
	. "github.com/chapzin/GoSped/Model"
	. "github.com/chapzin/GoSped/Utilidades"
)

var cofing2 = Config{}
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
		// var codigo string
		var chave string
		for _, linha := range conteudo {
			// fmt.Println(ind, linha)
			var reg0200 Reg0200

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
						chave = l[9]
					}

					if l[1] == "C170" {
						if len(l[3]) != 10 {
							fmt.Println(l[3] + ";" + chave)
						}

					}
					// }
				}
			}

		}
	}

}
