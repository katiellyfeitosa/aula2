package main

import (
	"fmt"
	"os"
)

func main() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Erro:", r)
		}
	}()

	lerArquivo("customers.txt")
}

func lerArquivo(nomeArquivo string) {

	data, err := os.ReadFile(nomeArquivo)

	if err != nil {
		panic("O arquivo indicado não foi encontrado ou está danificado ")
	}

	fmt.Printf("Arquivo %s lido com sucesso!\n", nomeArquivo)
	fmt.Println(string(data))
}
