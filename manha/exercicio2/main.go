package main

import "fmt"

func main() {
	var temperatura int
	var umidade float64
	var pressaoAtmosferica int

	temperatura, umidade, pressaoAtmosferica = 5, 81.3, 1016

	fmt.Println("A temperatura em São Paulo é:", temperatura, " graus, a umidade ", umidade, "e a pressão atmosferica", pressaoAtmosferica)

}
