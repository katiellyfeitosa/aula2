package main

import "fmt"

func main() {

	meses := []string{
		"Janeiro",
		"Fevereiro",
		"Março",
		"Abril",
		"Maio",
		"Junho",
		"Julho",
		"Agosto",
		"Setembro",
		"Outubro",
		"Novembro",
		"Dezembro",
	}

	var mes int = 2

	fmt.Println("Mês do ano: ", meses[mes-1])

}
