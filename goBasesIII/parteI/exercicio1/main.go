package main

import (
	"errors"
	"fmt"
	"os"
)

type produto struct {
	id         int
	quantidade int
	preco      float64
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func (p produto) gerarLinhaCSV() string {
	return fmt.Sprintf("%d,%d,%.2f\n", p.id, p.quantidade, p.preco)
}

func gerarCsv(caminho string, produtos []produto) error {

	if len(produtos) == 0 {
		return errors.New("Quantidade de produto inválida")
	}

	file, err := os.OpenFile(caminho, os.O_WRONLY|os.O_CREATE, 0600)

	check(err)

	p := produtos[0]
	file.WriteString("ID, Quantidade, Preco\n")
	for _, p = range produtos {
		_, err = file.WriteString(p.gerarLinhaCSV())
		check(err)
	}
	return nil
}

func main() {
	produtos := []produto{
		{
			id:         3,
			quantidade: 8,
			preco:      6.99,
		},
		{
			id:         2,
			quantidade: 20,
			preco:      12.99,
		},
		{
			id:         1,
			quantidade: 10,
			preco:      9.99,
		},
	}
	gerarCsv("./produtos.txt", produtos)
}
