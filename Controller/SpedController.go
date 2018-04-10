package Controller

import (
	"fmt"
	"strings"

	"github.com/chapzin/GoSped/Dao"

	"github.com/go-bongo/bongo"

	"github.com/chapzin/GoSped/ConfigTom"
	"github.com/chapzin/GoSped/Model"

	"github.com/chapzin/GoSped/Utilidades"
)

// SpedController : Responsavel por adicionar os dados do sped e validar o conteudo
type SpedController struct {
	linha string
}

var config2 ConfigTom.ConfigSped

func ProcessarPathSped(reg0000 Model.Reg0000) (pathArquivo string) {
	tipo := "sped"
	cnpj := reg0000.Cnpj
	ano := reg0000.DtIni[4:8]
	mes := reg0000.DtIni[2:4]
	path := "empresas"
	Utilidades.CriarEstruturaDePastas(path, cnpj, ano, mes, tipo)
	pathArquivo = path + "/" + cnpj + "/" + ano + "/" + mes + "/" + tipo + "/" + "Sped-" + mes + "-" + ano + ".txt"
	return pathArquivo
}

func (s *SpedController) addDB(arquivo string, conn *bongo.Connection) {
	var reg0000 Model.Reg0000
	var reg0100 Model.Reg0100
	var reg0150 Model.Reg0150
	var reg0190 Model.Reg0190
	var reg0200 Model.Reg0200
	var reg0220 Model.Reg0220
	var regC100 Model.RegC100
	var regC170 Model.RegC170
	var regC190 Model.RegC190
	var regC400 Model.RegC400
	var regC405 Model.RegC405
	var regC420 Model.RegC420
	var regC425 Model.RegC425
	var regC460 Model.RegC460
	var regC465 Model.RegC465
	var regC470 Model.RegC470
	var regC490 Model.RegC490
	var regH005 Model.RegH005
	var regH010 Model.RegH010
	conteudo, err := Utilidades.LerArquivoSped(arquivo)
	if err != nil {
		fmt.Println(err)
	}
	for _, linha := range conteudo {
		if linha != "" {
			if len(linha) > 0 {
				if linha[:1] == "|" {
					linhaSplit := strings.Split(linha, "|")
					if linhaSplit[1] == "0000" {
						reg0000 = Model.Reg0000{}
						reg0000.Populate(linhaSplit)
						err := conn.Collection(Dao.COLLECTION0000).Save(&reg0000)
						if err != nil {
							fmt.Println(err)
						}
						pathArquivo := ProcessarPathSped(reg0000)
						Utilidades.MoverArquivos(arquivo, pathArquivo)

					}
					if linhaSplit[1] == "0100" {
						reg0100 = Model.Reg0100{}
						reg0100.Populate(linhaSplit, reg0000)
						err := conn.Collection(Dao.COLLECTION0100).Save(&reg0100)
						if err != nil {
							fmt.Println(err)
						}
					}
					if linhaSplit[1] == "0150" {
						reg0150 = Model.Reg0150{}
						reg0150.Populate(linhaSplit, reg0000)
						err := conn.Collection(Dao.COLLECTION0150).Save(&reg0150)
						if err != nil {
							fmt.Println(err)
						}
					}
					if linhaSplit[1] == "0190" {
						reg0190 = Model.Reg0190{}
						reg0190.Populate(linhaSplit, reg0000)
						err := conn.Collection(Dao.COLLECTION0190).Save(&reg0190)
						if err != nil {
							fmt.Println(err)
						}
					}
					if linhaSplit[1] == "0200" {
						reg0200 = Model.Reg0200{}
						reg0200.Populate(linhaSplit, reg0000)
						err := conn.Collection(Dao.COLLECTION0200).Save(&reg0200)
						if err != nil {
							fmt.Println(err)
						}
					}
					if linhaSplit[1] == "0220" {
						reg0220 = Model.Reg0220{}
						reg0220.Populate(linhaSplit, reg0000, reg0200)
						err := conn.Collection(Dao.COLLECTION0220).Save(&reg0220)
						if err != nil {
							fmt.Println(err)
						}
					}
					if linhaSplit[1] == "C100" {
						regC100 = Model.RegC100{}
						regC100.Populate(linhaSplit, reg0000)
						err := conn.Collection(Dao.COLLECTIONC100).Save(&regC100)
						if err != nil {
							fmt.Println(err)
						}
					}
					if linhaSplit[1] == "C170" {
						regC170 = Model.RegC170{}
						regC170.Populate(linhaSplit, reg0000, regC100)
						err := conn.Collection(Dao.COLLECTIONC170).Save(&regC170)
						if err != nil {
							fmt.Println(err)
						}

					}
					if linhaSplit[1] == "C190" {
						regC190 = Model.RegC190{}
						regC190.Populate(linhaSplit, reg0000)
						err := conn.Collection(Dao.COLLECTIONC190).Save(&regC190)
						if err != nil {
							fmt.Println(err)
						}
					}
					if linhaSplit[1] == "C400" {
						regC400 = Model.RegC400{}
						regC400.Populate(linhaSplit, reg0000)
						err := conn.Collection(Dao.COLLECTIONC400).Save(&regC400)
						if err != nil {
							fmt.Println(err)
						}
					}
					if linhaSplit[1] == "C405" {
						regC405 = Model.RegC405{}
						regC405.Populate(linhaSplit, reg0000)
						err := conn.Collection(Dao.COLLECTIONC405).Save(&regC405)
						if err != nil {
							fmt.Println(err)
						}
					}
					if linhaSplit[1] == "C420" {
						regC420 = Model.RegC420{}
						regC420.Populate(linhaSplit, reg0000)
						err := conn.Collection(Dao.COLLECTIONC420).Save(&regC420)
						if err != nil {
							fmt.Println(err)
						}
					}
					if linhaSplit[1] == "C425" {
						regC425 = Model.RegC425{}
						regC425.Populate(linhaSplit, reg0000, regC405)
						err := conn.Collection(Dao.COLLECTIONC425).Save(&regC425)
						if err != nil {
							fmt.Println(err)
						}
					}
					if linhaSplit[1] == "C460" {
						regC460 = Model.RegC460{}
						regC460.Populate(linhaSplit, reg0000)
						err := conn.Collection(Dao.COLLECTIONC460).Save(&regC460)
						if err != nil {
							fmt.Println(err)
						}
					}
					if linhaSplit[1] == "C465" {
						regC465 = Model.RegC465{}
						regC465.Populate(linhaSplit, reg0000)
						err := conn.Collection(Dao.COLLECTIONC465).Save(&regC465)
						if err != nil {
							fmt.Println(err)
						}
					}
					if linhaSplit[1] == "C470" {
						regC470 = Model.RegC470{}
						regC470.Populate(linhaSplit, reg0000)
						err := conn.Collection(Dao.COLLECTIONC470).Save(&regC470)
						if err != nil {
							fmt.Println(err)
						}
					}
					if linhaSplit[1] == "C490" {
						regC490 = Model.RegC490{}
						regC490.Populate(linhaSplit, reg0000)
						err := conn.Collection(Dao.COLLECTIONC490).Save(&regC490)
						if err != nil {
							fmt.Println(err)
						}
					}
					if linhaSplit[1] == "H005" {
						regH005 = Model.RegH005{}
						regH005.Populate(linhaSplit, reg0000)
						err := conn.Collection(Dao.COLLECTIONH005).Save(&regH005)
						if err != nil {
							fmt.Println(err)
						}
					}
					if linhaSplit[1] == "H010" {
						regH010 = Model.RegH010{}
						regH010.Populate(linhaSplit, reg0000, regH005)
						err := conn.Collection(Dao.COLLECTIONH010).Save(&regH010)
						if err != nil {
							fmt.Println(err)
						}
					}
				}
			}
		}
	}
}

func (s *SpedController) validacoes(arquivo string) {
	conteudo, err := Utilidades.LerArquivoSped(arquivo)
	if err != nil {
		fmt.Println(err)
	}
	for _, linha := range conteudo {
		if linha != "" {
			if len(linha) > 0 {
				if linha[:1] == "|" {
					linhaSplit := strings.Split(linha, "|")
					if linhaSplit[1] == "0000" {

					}
					if linhaSplit[1] == "0100" {

					}
					if linhaSplit[1] == "0150" {

					}
					if linhaSplit[1] == "0190" {

					}
					if linhaSplit[1] == "0200" {

					}
					if linhaSplit[1] == "0220" {

					}
					if linhaSplit[1] == "C100" {

					}
					if linhaSplit[1] == "C170" {

					}
					if linhaSplit[1] == "C190" {

					}
					if linhaSplit[1] == "C400" {

					}
					if linhaSplit[1] == "C405" {

					}
					if linhaSplit[1] == "C420" {

					}
					if linhaSplit[1] == "C425" {

					}
					if linhaSplit[1] == "C460" {

					}
					if linhaSplit[1] == "C465" {

					}
					if linhaSplit[1] == "C470" {

					}
					if linhaSplit[1] == "C490" {

					}
					if linhaSplit[1] == "H005" {

					}
					if linhaSplit[1] == "H010" {

					}
				}
			}
		}
	}
}
