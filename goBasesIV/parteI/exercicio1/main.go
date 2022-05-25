package main

import (
	"fmt"
)

type error interface {
	Error() string
}

type erroPersonalizado struct {
	msg string
}

func (e *erroPersonalizado) Error() string {
	return fmt.Sprintf("%s", e.msg)
}

func verificaSalario(salario int) error {
	if salario < 15000 {
		return &erroPersonalizado{
			msg: "erro: O salário digitado não está dentro do valor mínimo",
		}
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
	salario := 1000

	err := verificaSalario(salario)
	checkError(err)

}
