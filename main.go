package main

import (
	"github.com/chapzin/GoSped/Controller"
	"github.com/chapzin/GoSped/Utilidades"
)

func main() {
	var importar Controller.ImportController

	arquivos, _ := Utilidades.ListFiles("./importar/")
	for _, arq := range arquivos {
		importar.Importar(arq)
	}

}
