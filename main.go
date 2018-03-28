package main

import (
	"github.com/chapzin/GoSped/ConfigTom"
	"github.com/chapzin/GoSped/Controller"
	"github.com/chapzin/GoSped/Utilidades"
)

var cofing2 = ConfigTom.ConfigSped{}

func init() {
	cofing2.Read()
}

func main() {
	var importar Controller.ImportController

	arquivos, _ := Utilidades.ListFiles(cofing2.PathImport)
	for _, arq := range arquivos {
		importar.Importar(arq)
	}

}
