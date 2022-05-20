package main

import "fmt"


func Calc(base int, Horas int, taxa float32, minHoras int) float32{
salario := float32(base) * float32(Horas)
	if Horas >  minHoras {
	salario += salario * taxa
	}
return salario
}

func Salario(funcionarios []map[string]interface{}) float32 {

var salario float32

for _, v := range funcionarios {
fmt.Println("funcionario: ", v["Nome"])
categoria := v["Categoria"].(string)
Horas := v["Horas"].(int)



	if categoria == "C" {
	
	salario = float32(Horas * 1000)
	fmt.Println("salario: ", salario)
	}
	if categoria == "B" {
		salario = Calc(1500, Horas, 0.20, 160)
		fmt.Println("salario: ", salario)
	}
	if categoria == "A" {
		salario = Calc(3000, Horas, 0.50, 160)
		fmt.Println("salario: ", salario)
	}

}

return salario
}

func main() {

funcionarios := []map[string]interface{}{
{
"Nome": "João",
"Categoria": "C",
"Horas": 162,
},
{
"Nome": "Pedro",
"Categoria": "B",
"Horas": 176,
},
{
"Nome": "Maria",
"Categoria": "A",
"Horas": 172,
},
}

Salario(funcionarios)

}