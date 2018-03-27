package Utilidades

import (
	"bufio"
	"os"
	"path/filepath"
)

// ReadFile : Ler o arquivo texto e retorna slice de string
func ReadFile(caminhoArquivo string) ([]string, error) {
	arquivo, err := os.Open(caminhoArquivo)
	if err != nil {
		return nil, err
	}

	defer arquivo.Close()

	var linhas []string
	scanner := bufio.NewScanner(arquivo)
	for scanner.Scan() {
		linhas = append(linhas, scanner.Text())
	}

	return linhas, scanner.Err()
}

// ListarArquivos :  Apenas uma conversao de nome do glob para ListarArquivos
func ListFiles(caminho string) ([]string, error) {
	files, err := filepath.Glob(caminho + "*.txt")
	if err != nil {
		return nil, err
	}
	return files, err
}
