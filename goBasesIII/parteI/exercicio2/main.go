package main

import (
	"fmt"
	"os"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	data, err := os.ReadFile("./produtos.txt")
	check(err)
	fmt.Println(string(data))
}
