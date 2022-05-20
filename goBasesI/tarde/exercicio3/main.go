package main

import "fmt"

func main() {

	meses := []string{"janeiro", "fevereiro", "março", "abril", "maio", "junho", "julho", "agosto", "setembro", "outubro", "novembro", "dezembro"}

	var mes int = 2

	fmt.Println("Mês do ano: ", meses[mes-1])

}
