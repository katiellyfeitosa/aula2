package main

import "fmt"

func main() {
var employees = map[string]int{"Benjamin": 20, "Manuel": 26, "Brenda": 19, "Dario": 44, "Pedro": 30}

fmt.Println("Idade de Benjamin", employees["Benjamin"])

var contador int = 0

for _, value := range employees {
	if value > 21 {
		contador++
	}
}

fmt.Println("Valor de contador: ", contador)

employees["Frederico"] = 25

delete(employees, "Pedro")

fmt.Println(employees)

}