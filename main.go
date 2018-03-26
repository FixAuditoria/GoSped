package main

import (
	"fmt"
	"log"
	"strings"

	"github.com/chapzin/IiniciandoGo/Model"
	"github.com/chapzin/IiniciandoGo/Utilidades"
)

func main() {
	arquivos, _ := Utilidades.ListarArquivos("./speds/")
	for _, arq := range arquivos {
		conteudo, err := Utilidades.LerTexto(arq)
		if err != nil {
			log.Fatalf("Erro:", err)
		}
		var reg0000 string
		// var codigo string
		var chave string
		for _, linha := range conteudo {
			// fmt.Println(ind, linha)
			var reg0200 Model.Reg0200

			if len(linha) > 0 {
				if linha[:1] == "|" {
					l := strings.Split(linha, "|")
					if l[1] == "0000" {
						reg0000 = linha
						fmt.Println(reg0000)
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
