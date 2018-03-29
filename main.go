package main

import (
	"fmt"

	"github.com/chapzin/GoSped/ConfigTom"
	"github.com/chapzin/GoSped/Dao"
)

var cofing2 = ConfigTom.ConfigSped{}
var regC170Dao Dao.RegC170Dao

func init() {
	cofing2.Read()
	regC170Dao.Server = cofing2.Server
	regC170Dao.Database = cofing2.Database
	regC170Dao.Connect()
}

func main() {
	regc170s, err := regC170Dao.FindByCnpj("07965203000161")
	if err != nil {
		fmt.Println(err)
	}
	for _, regc170 := range regc170s {
		if len(regc170.CodItem) != 10 {
			fmt.Println(regc170.CodItem)
		}

	}
	// ----------------- Importa todos arquivos da pasta que fica no config.toml ------------------------
	// var importar Controller.ImportController

	// arquivos, _ := Utilidades.ListFiles(cofing2.PathImport)
	// for _, arq := range arquivos {
	// 	importar.Importar(arq)
	// }

}
