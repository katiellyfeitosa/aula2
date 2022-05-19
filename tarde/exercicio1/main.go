package main

import (
"fmt"
"strings"
)

func main() {

var palavra string = "hello"

fmt.Println("Quantidade de letras: ", len(palavra))

str := strings.Split(palavra,"")

for i := 0; i < len(str); i++ {
	fmt.Println(str[i])
}

}