package main

import (
	"fmt"
)

func main() {

	var palavra string = "hello"

	fmt.Println("Quantidade de letras: ", len(palavra))

	for _, i := range palavra {
		fmt.Println(string(i))
	}

}
