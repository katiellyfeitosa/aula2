package main

import (
	"fmt"
)

func verificaSalario(salario int) error {
	if salario < 15000 {
		return fmt.Errorf("O mínimo tributável é 15.000 e o salário informado é: %d", salario)
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
	salario := 8000

	err := verificaSalario(salario)
	checkError(err)

}
