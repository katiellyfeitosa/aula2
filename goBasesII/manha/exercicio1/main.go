package main

import (
	"fmt"
)

func main() {
	funcionario1 := calcImposto(50000, 17)
	funcionario2 := calcImposto(150000, 10)

	fmt.Println("O funcioanrio 1 vai pagar", funcionario1, "de imposto.")
	fmt.Println("O funcioanrio 2 vai pagar", funcionario2, "de imposto.")

}

func calcImposto(salario float32, taxaImposto float32) float32 {

	return (salario * taxaImposto) / 100
}
