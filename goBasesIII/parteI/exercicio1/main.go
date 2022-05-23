package main

import (
	"fmt"
	"os"
)

func main() {
	produtos := []byte("ID 	NOME  PRECO  QUANTIDADE\n01, Candida, 20.0, 20;\n02, Desinfetante, 15.0, 6;\n03, Amaciante, 16.0, 50;\n04, Limpa pedra, 90.0, 5;")

	err := os.WriteFile("./produtos.txt", produtos, 0644)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Arquivo gerado com sucesso!")
	}

}
