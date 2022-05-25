package main

import (
	"errors"
	"fmt"
)

type Usuario struct {
	Nome      string
	Sobrenome string
	email     string
	Produtos  []Produto
}

type Produto struct {
	Nome       string
	preco      float64
	quantidade int
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	var usuario1 *Usuario
	var produto1 *Produto

	produto1 = novoProduto("Caixa de som", 499.99)

	usuario1 = &Usuario{
		Nome:      "Katielly",
		Sobrenome: "Rezende",
		email:     "teste@teste.com",
	}

	fmt.Println("Usuario:", usuario1)

	err := usuario1.adicionarProduto(produto1, 10)
	check(err)
	fmt.Println("Usuario com um produto:", usuario1)

	usuario1.deletarProdutos()
	fmt.Println("Usuario após a exclusao dos seus produtos", usuario1)

}

func novoProduto(nome string, preco float64) *Produto {
	return &Produto{
		Nome:  nome,
		preco: preco,
	}
}

func (u *Usuario) adicionarProduto(p *Produto, quantidade int) error {
	if quantidade == 0 {
		return errors.New("Quantidade de produto inválida")
	}

	p.quantidade = quantidade
	u.Produtos = append(u.Produtos, *p)

	return nil
}

func (u *Usuario) deletarProdutos() {
	u.Produtos = nil
}
