package main

import "fmt"

type Produto struct {
	Nome       string
	preco      float64
	quantidade int
}

type Servico struct {
	Nome           string
	preco          float64
	minTrabalhados float64
}

type Manutencao struct {
	Nome  string
	preco float64
}

func main() {
	var arrayProd []Produto

	prod1 := Produto{
		Nome:       "Produto 1",
		preco:      50,
		quantidade: 2,
	}

	prod2 := Produto{
		Nome:       "Produto 2",
		preco:      55.50,
		quantidade: 3,
	}
	arrayProd = append(arrayProd, prod1)
	arrayProd = append(arrayProd, prod2)

	var arrayServico []Servico

	serv1 := Servico{
		Nome:           "Servico 1",
		preco:          1,
		minTrabalhados: 25,
	}

	serv2 := Servico{
		Nome:           "Servico 2",
		preco:          1,
		minTrabalhados: 90,
	}

	arrayServico = append(arrayServico, serv1)
	arrayServico = append(arrayServico, serv2)

	var arrayManutencao []Manutencao

	manutencao1 := Manutencao{
		Nome:  "Manutencao 1",
		preco: 10,
	}

	manutencao2 := Manutencao{
		Nome:  "Mnautencao 2",
		preco: 10,
	}

	arrayManutencao = append(arrayManutencao, manutencao1)
	arrayManutencao = append(arrayManutencao, manutencao2)

	p := make(chan float64)
	s := make(chan float64)
	m := make(chan float64)

	go somarServico(arrayServico, s)
	go somarManutencao(arrayManutencao, m)
	go somarProdutos(arrayProd, p)

	resultadoProduto := <-p
	resultadoServico := <-s
	resultadoManutencao := <-m

	total := resultadoProduto + resultadoServico + resultadoManutencao
	fmt.Printf("A soma dos produtos, servicos e manutecoes é R$%.2f\n", total)
}

func somarProdutos(produtos []Produto, p chan float64) {
	var precoTotal float64
	for i, _ := range produtos {
		precoTotal += produtos[i].preco * float64(produtos[i].quantidade)
	}
	fmt.Println("Produtos", precoTotal)
	p <- precoTotal
}

func somarServico(servicos []Servico, s chan float64) {
	var somaServicos float64
	for i, _ := range servicos {
		if servicos[i].minTrabalhados < 30 {
			somaServicos += servicos[i].preco * 30
		} else {
			somaServicos += servicos[i].preco * servicos[i].minTrabalhados
		}
	}
	fmt.Println("Servicos", somaServicos)
	s <- somaServicos
}

func somarManutencao(manutecoes []Manutencao, m chan float64) {
	var somaManutencao float64
	for i, _ := range manutecoes {
		somaManutencao += manutecoes[i].preco
	}
	fmt.Println("Manutencao", somaManutencao)
	m <- somaManutencao

}
