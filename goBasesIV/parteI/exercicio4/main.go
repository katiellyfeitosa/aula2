// Exercício 4 - Imposto sobre o salário #4
// Vamos fazer com que nosso programa seja um pouco mais complexo e útil.
// 1. Desenvolva as funções necessárias para permitir que a empresa calcule:
// a) Salário mensal de um funcionário segundo a quantidade de horas trabalhadas.
// - A função receberá as horas trabalhadas no mês e o valor da hora como
// parâmetro.
// - Esta função deve retornar mais de um valor (salário calculado e erro).
// - No caso de o salário mensal ser igual ou superior a R$ 20.000, 10% devem ser
// descontados como imposto.
// Se o número de horas mensais inseridas for menor que 80 ou um número
// negativo, a função deverá retornar um erro. Deve indicar "erro: o trabalhador
// não pode ter trabalhado menos de 80 horas por mês".

// 2. Desenvolva o código necessário para cumprir as funcionalidades solicitadas, usando
// “errors.New()”, “fmt.Errorf()” e “errors.Unwrap()”. Não esqueça de realizar as validações dos
// retornos de erro em sua função “main()”.

package main

import (
	"errors"
	"fmt"
)

type funcionario struct {
	Nome           string
	horasTrabalhas int
	valorHora      float64
}

func (f funcionario) calcSalario() (float64, error) {

	if f.horasTrabalhas == 0 {
		return 0, errors.New("o valor de horas trabalhadas não pode ser 0")
	}

	if f.horasTrabalhas < 80 {
		return 0, fmt.Errorf("O trabalhador %s não pode ter trabalhado menos de 80 horas por mês", f.Nome)
	}

	salario := float64(f.horasTrabalhas) * f.valorHora

	salarioCalc := calcImposto(salario)
	return salarioCalc, nil
}

func calcImposto(salario float64) float64 {
	if salario >= 20000 {
		salario -= salario * 0.10
		fmt.Printf("Imposto aplicado")
	}
	return salario
}

func checkError(e error) {
	if e != nil {
		fmt.Println(e)
	}
}

func main() {
	funcionario1 := funcionario{
		Nome:           "Katielly",
		horasTrabalhas: 20,
		valorHora:      500,
	}

	salario, err := funcionario1.calcSalario()

	checkError(err)
	fmt.Printf("O salario do funcionario %s é R$%.2f\n", funcionario1.Nome, salario)
}
