package Utilidades

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"path/filepath"
)

// LerArquivoSped : Ler o arquivo texto e retorna slice de string
func LerArquivoSped(caminhoArquivo string) ([]string, error) {
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

func DiretorioVazio(path string) {
	diretorio, err := os.Open(path)
	if err != nil {
		fmt.Println(err.Error())
	}
	defer diretorio.Close()
	_, err = diretorio.Readdirnames(1)
	if err == io.EOF {
		os.Remove(path)
	}
}

// ListarArquivosV2 :  Apenas uma conversao de nome do glob para ListarArquivos
func ListarArquivosV2(caminho string) (files []string, err error) {
	err = filepath.Walk(caminho, func(path string, info os.FileInfo, err error) error {
		if !info.IsDir() {
			files = append(files, path)
		} else {
			DiretorioVazio(path)
		}
		return nil
	})
	if err != nil {
		return nil, err
	}
	return files, nil
}

// CriarUmDiretorio : Funcao responsavel por criar um diretorio caso ele nao exista
func CriarUmDiretorio(path string) {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		os.Mkdir(path, 0755)
	}
}

// CriarEstruturaDePastas : Funcao responsavel por criar uma estrutura grande de pastas para o armazenamento dos xmls
func CriarEstruturaDePastas(path string, cnpj string, ano string, mes string, tipo string) (pathFinal string) {
	pathcnpj := path + "/" + cnpj
	fmt.Println(pathcnpj)
	pathano := path + "/" + cnpj + "/" + ano
	pathmes := path + "/" + cnpj + "/" + ano + "/" + mes
	pathtipo := path + "/" + cnpj + "/" + ano + "/" + mes + "/" + tipo

	CriarUmDiretorio(pathcnpj)
	CriarUmDiretorio(pathano)
	CriarUmDiretorio(pathmes)
	CriarUmDiretorio(pathtipo)
	return pathtipo
}

// MoverXml : Funcao responsavel por levar o xml do pathInicial para o pathFinal
func MoverArquivos(pathInicial string, pathFinal string) {
	err := os.Rename(pathInicial, pathFinal)
	if err != nil {
		fmt.Println(err)
	}
}
