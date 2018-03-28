package Controller

import (
	"log"
	"strings"

	"github.com/chapzin/GoSped/Utilidades"
)

type SpedController struct {
	linha string
}

func (s *SpedController) addMongo(arquivo string) {
	conteudo, err := Utilidades.ReadFile(arquivo)
	if err != nil {
		log.Fatalf("Erro:", err)
	}
	for _, linha := range conteudo {
		if linha != "" {
			if len(linha) > 0 {
				if linha[:1] == "|" {
					l := strings.Split(linha, "|")
					if l[1] == "0000" {

					}
					if l[1] == "0100" {

					}
					if l[1] == "0150" {

					}
					if l[1] == "0190" {

					}
					if l[1] == "0200" {

					}
					if l[1] == "0220" {

					}
					if l[1] == "C100" {

					}
					if l[1] == "C170" {

					}
					if l[1] == "C190" {

					}
					if l[1] == "C400" {

					}
					if l[1] == "C405" {

					}
					if l[1] == "C420" {

					}
					if l[1] == "C425" {

					}
					if l[1] == "C460" {

					}
					if l[1] == "C465" {

					}
					if l[1] == "C470" {

					}
					if l[1] == "C490" {

					}
					if l[1] == "H005" {

					}
					if l[1] == "H010" {

					}
				}
			}
		}
	}
}

func (s *SpedController) validacoes(arquivo string) {
	conteudo, err := Utilidades.ReadFile(arquivo)
	if err != nil {
		log.Fatalf("Erro:", err)
	}
	for _, linha := range conteudo {
		if linha != "" {
			if len(linha) > 0 {
				if linha[:1] == "|" {
					l := strings.Split(linha, "|")
					if l[1] == "0000" {

					}
					if l[1] == "0100" {

					}
					if l[1] == "0150" {

					}
					if l[1] == "0190" {

					}
					if l[1] == "0200" {

					}
					if l[1] == "0220" {

					}
					if l[1] == "C100" {

					}
					if l[1] == "C170" {

					}
					if l[1] == "C190" {

					}
					if l[1] == "C400" {

					}
					if l[1] == "C405" {

					}
					if l[1] == "C420" {

					}
					if l[1] == "C425" {

					}
					if l[1] == "C460" {

					}
					if l[1] == "C465" {

					}
					if l[1] == "C470" {

					}
					if l[1] == "C490" {

					}
					if l[1] == "H005" {

					}
					if l[1] == "H010" {

					}
				}
			}
		}
	}
}
