package main

import (
	"fmt"
	"time"
)

func main() {
	// var cliente = map[string]interface{}{"nome": "Joao", "idade": 20, "empregado": true, "dataDeAdmissao": "2022-05-01", "salario": 100.000}
	var (
		nome           = "Joao"
		idade          = 20
		empregado      = true
		dataDeAdmissao = "2022-05-01T00:00:00Z"
		salario        = 100.000
	)

	dataAgora := time.Now()
	dataFormatada, _ := time.Parse(time.RFC3339, dataDeAdmissao)
	tempoTrabalho := dataAgora.Sub(dataFormatada)
	// tempo := tempoTrabalho / 24
	fmt.Println(tempoTrabalho)

	if idade > 22 {
		if empregado == true {
			if salario > 100.000 {
				fmt.Println("O cliente", nome, "está elegivel, tem direito ao emprestimo com taxas de juros menores.")
			} else {
				fmt.Println("O cliente", nome, "está elegivel para fazer emprestimos.")
			}
		} else {
			fmt.Println("O cliente", nome, "não está elegivel, não fornecemos emprestimos para desempregados.")
		}
	} else {
		fmt.Println("O cliente", nome, "não está elegivel, não fornecemos emprestimos para menores de 21 anos.")
	}
}
