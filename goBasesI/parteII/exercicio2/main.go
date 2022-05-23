package main

import (
	"fmt"
)

func main() {
	clientes := []map[string]interface{}{
		{
			"Nome":      "João",
			"Idade":     21,
			"Atividade": 1,
			"Salário":   100000,
		},
		{
			"Nome":      "José",
			"Idade":     25,
			"Atividade": 1,
			"Salário":   100000,
		},
		{
			"Nome":      "Carlos",
			"Idade":     27,
			"Atividade": 2,
			"Salário":   99000,
		},
		{
			"Nome":      "Antônio",
			"Idade":     35,
			"Atividade": 5,
			"Salário":   150000,
		},
	}

	for _, c := range clientes {
		nome := c["Nome"]
		fmt.Println("Cliente:", nome)
		idade := c["Idade"].(int) //cast para int
		if idade <= 22 {
			fmt.Println("Não possui empréstimo disponível (idade insuficiente)")
			continue
		}

		atividade := c["Atividade"].(int) //cast para int
		if atividade <= 1 {
			fmt.Println("Não possui empréstimo disponível (atividade insuficiente)")
			continue
		}

		salario := c["Salário"].(int) //cast para int
		if salario > 100000 {
			fmt.Println("Possui empréstimo disponível sem juros")
		} else {
			fmt.Println("Possui empréstimo disponível com juros")
		}
	}
}
