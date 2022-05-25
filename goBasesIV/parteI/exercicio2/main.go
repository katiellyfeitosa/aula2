package main

import (
	"errors"
	"fmt"
)

func verificaSalario(salario int) error {
	if salario < 15000 {
		return errors.New("O salário digitado não está dentro do valor mínimo")
	}
	fmt.Println("Necessário pagamento de imposto")
	return nil
}

func checkError(e error) {
	if e != nil {
		fmt.Println(e)
	}
}

func main() {
	salario := 9000

	err := verificaSalario(salario)
	checkError(err)

}
