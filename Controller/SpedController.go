package Controller

import (
	"fmt"
	"log"
	"strings"

	"github.com/chapzin/GoSped/Dao"
	"github.com/chapzin/GoSped/Model"

	"github.com/chapzin/GoSped/ConfigTom"

	"github.com/chapzin/GoSped/Utilidades"
)

type SpedController struct {
	linha string
}

var config2 ConfigTom.ConfigSped
var reg0000Dao Dao.Reg0000Dao
var reg0100Dao Dao.Reg0100Dao
var reg0150Dao Dao.Reg0150Dao
var reg0190Dao Dao.Reg0190Dao
var reg0200Dao Dao.Reg0200Dao
var reg0220Dao Dao.Reg0220Dao
var regC100Dao Dao.RegC100Dao
var regC170Dao Dao.RegC170Dao
var regC190Dao Dao.RegC190Dao
var regC400Dao Dao.RegC400Dao
var regC405Dao Dao.RegC405Dao
var regC420Dao Dao.RegC420Dao
var regC425Dao Dao.RegC425Dao
var regC460Dao Dao.RegC460Dao
var regC465Dao Dao.RegC465Dao
var regC470Dao Dao.RegC470Dao
var regC490Dao Dao.RegC490Dao
var regH005Dao Dao.RegH005Dao
var regH010Dao Dao.RegH010Dao

func init() {
	config2.Read()
	reg0000Dao.Server = config2.Server
	reg0000Dao.Database = config2.Database
	reg0000Dao.Connect()

	reg0100Dao.Server = config2.Server
	reg0100Dao.Database = config2.Database
	reg0100Dao.Connect()

	reg0150Dao.Server = config2.Server
	reg0150Dao.Database = config2.Database
	reg0150Dao.Connect()

	reg0190Dao.Server = config2.Server
	reg0190Dao.Database = config2.Database
	reg0190Dao.Connect()

	reg0200Dao.Server = config2.Server
	reg0200Dao.Database = config2.Database
	reg0200Dao.Connect()

	reg0220Dao.Server = config2.Server
	reg0220Dao.Database = config2.Database
	reg0220Dao.Connect()

	regC100Dao.Server = config2.Server
	regC100Dao.Database = config2.Database
	regC100Dao.Connect()

	regC170Dao.Server = config2.Server
	regC170Dao.Database = config2.Database
	regC170Dao.Connect()

	regC190Dao.Server = config2.Server
	regC190Dao.Database = config2.Database
	regC190Dao.Connect()

	regC400Dao.Server = config2.Server
	regC400Dao.Database = config2.Database
	regC400Dao.Connect()

	regC405Dao.Server = config2.Server
	regC405Dao.Database = config2.Database
	regC405Dao.Connect()

	regC420Dao.Server = config2.Server
	regC420Dao.Database = config2.Database
	regC420Dao.Connect()

	regC425Dao.Server = config2.Server
	regC425Dao.Database = config2.Database
	regC425Dao.Connect()

	regC460Dao.Server = config2.Server
	regC460Dao.Database = config2.Database
	regC460Dao.Connect()

	regC465Dao.Server = config2.Server
	regC465Dao.Database = config2.Database
	regC465Dao.Connect()

	regC470Dao.Server = config2.Server
	regC470Dao.Database = config2.Database
	regC470Dao.Connect()

	regC490Dao.Server = config2.Server
	regC490Dao.Database = config2.Database
	regC490Dao.Connect()

	regH005Dao.Server = config2.Server
	regH005Dao.Database = config2.Database
	regH005Dao.Connect()

	regH010Dao.Server = config2.Server
	regH010Dao.Database = config2.Database
	regH010Dao.Connect()
}

func (s *SpedController) addMongo(arquivo string) {
	conteudo, err := Utilidades.ReadFile(arquivo)
	if err != nil {
		log.Fatalf("Erro:", err)
	}
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

	for _, linha := range conteudo {
		if linha != "" {
			if len(linha) > 0 {
				if linha[:1] == "|" {
					l := strings.Split(linha, "|")
					if l[1] == "0000" {
						reg0000.Populate(l)
						err := reg0000Dao.Insert(reg0000)
						if err != nil {
							fmt.Println(err)
						}
					}
					if l[1] == "0100" {
						reg0100.Populate(l, reg0000)
						err := reg0100Dao.Insert(reg0100)
						if err != nil {
							fmt.Println(err)
						}
					}
					if l[1] == "0150" {
						reg0150.Populate(l, reg0000)
						err := reg0150Dao.Insert(reg0150)
						if err != nil {
							fmt.Println(err)
						}

					}
					if l[1] == "0190" {
						reg0190.Populate(l, reg0000)
						err := reg0190Dao.Insert(reg0190)
						if err != nil {
							fmt.Println(err)
						}
					}
					if l[1] == "0200" {
						reg0200.Populate(l, reg0000)
						err := reg0200Dao.Insert(reg0200)
						if err != nil {
							fmt.Println(err)
						}

					}
					if l[1] == "0220" {
						reg0220.Populate(l, reg0000, reg0200)
						err := reg0220Dao.Insert(reg0220)
						if err != nil {
							fmt.Println(err)
						}

					}
					if l[1] == "C100" {
						regC100.Populate(l, reg0000)
						err := regC100Dao.Insert(regC100)
						if err != nil {
							fmt.Println(err)
						}

					}
					if l[1] == "C170" {
						regC170.Populate(l, reg0000, regC100)
						err := regC170Dao.Insert(regC170)
						if err != nil {
							fmt.Println(err)
						}

					}
					if l[1] == "C190" {
						regC190.Populate(l, reg0000)
						err := regC190Dao.Insert(regC190)
						if err != nil {
							fmt.Println(err)
						}
					}
					if l[1] == "C400" {
						regC400.Populate(l, reg0000)
						err := regC400Dao.Insert(regC400)
						if err != nil {
							fmt.Println(err)
						}
					}
					if l[1] == "C405" {
						regC405.Populate(l, reg0000)
						err := regC405Dao.Insert(regC405)
						if err != nil {
							fmt.Println(err)
						}

					}
					if l[1] == "C420" {
						regC420.Populate(l, reg0000)
						err := regC420Dao.Insert(regC420)
						if err != nil {
							fmt.Println(err)
						}
					}
					if l[1] == "C425" {
						regC425.Populate(l, reg0000, regC405)
						err := regC425Dao.Insert(regC425)
						if err != nil {
							fmt.Println(err)
						}

					}
					if l[1] == "C460" {
						regC460.Populate(l, reg0000)
						err := regC460Dao.Insert(regC460)
						if err != nil {
							fmt.Println(err)
						}

					}
					if l[1] == "C465" {
						regC465.Populate(l, reg0000)
						err := regC465Dao.Insert(regC465)
						if err != nil {
							fmt.Println(err)
						}
					}
					if l[1] == "C470" {
						regC470.Populate(l, reg0000)
						err := regC470Dao.Insert(regC470)
						if err != nil {
							fmt.Println(err)
						}

					}
					if l[1] == "C490" {
						regC490.Populate(l, reg0000)
						err := regC490Dao.Insert(regC490)
						if err != nil {
							fmt.Println(err)
						}

					}
					if l[1] == "H005" {
						regH005.Populate(l, reg0000)
						err := regH005Dao.Insert(regH005)
						if err != nil {
							fmt.Println(err)
						}

					}
					if l[1] == "H010" {
						regH010.Populate(l, reg0000, regH005)
						err := regH010Dao.Insert(regH010)
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
