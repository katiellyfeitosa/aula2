package main

import (
	"fmt"
	"sort"
)

func main() {
	const (
		minimum = "minimum"
		average = "average"
		maximum = "maximum"
	)

	minhaFunc := operacao("adicao")
	averageFunc := operacao("average")
	maxFunc := operacao("maximum")

	minValue := minhaFunc([]int{2, 3, 3, 4, 10, 2, 4, 5})
	averageValue := averageFunc([]int{2, 3, 3, 4, 1, 2, 4, 5})
	maxValue := maxFunc([]int{2, 3, 3, 4, 15, 2, 4, 5})

	fmt.Println("Menor nota:", minValue)
	fmt.Println("Media:", averageValue)
	fmt.Println("Maximo:", maxValue)
}

func operacao(operador string) func(valor []int) int {
	switch operador {
	case "minimum":
		return opMinimum
	case "average":
		return opAverage
	case "maximum":
		return opMaximum
	}
	fmt.Println("Tipo de operacao indefinida")
	return nil
}

func opMinimum(valor []int) int {
	sort.Ints(valor)

	return valor[0]
}

func opAverage(valor []int) int {
	var soma int

	for _, v := range valor {
		soma += v
	}

	result := soma / len(valor)
	return result
}

func opMaximum(valor []int) int {
	sort.Ints(valor)

	l := len(valor)
	if l == 0 {
		return 0
	}

	return valor[l-1]
}
