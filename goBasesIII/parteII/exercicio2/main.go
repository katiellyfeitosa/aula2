// - Deletar produtos: recebe um usuário, apaga os produtos do usuário. FALTA
package main

import "fmt"

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

	usuario1.adicionarProduto(produto1, 10)

	fmt.Println("Usuario com um produto:", usuario1)
}

func novoProduto(nome string, preco float64) *Produto {
	return &Produto{
		Nome:  nome,
		preco: preco,
	}
}

func (u *Usuario) adicionarProduto(p *Produto, quantidade int) {
	p.quantidade = quantidade
	u.Produtos = append(u.Produtos, *p)
}

func (u *Usuario) deletarProduto(p *Produto) {
}
