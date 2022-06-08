// Exercício 3 - Listar Entidade

// Já tendo criado e testado nossa API que nos recebe, geramos uma rota que retorna uma lista
// do tema escolhido.
// 1. Dentro do "main.go", crie uma estrutura de acordo com o tema com os campos
// correspondentes.
// 2. Crie um endpoint cujo caminho é /thematic (plural). Exemplo: “/products”
// 3. Crie uma handler para o endpoint chamada "GetAll".
// 4. Crie um slice da estrutura e retorne-a por meio de nosso endpoint.
package main

import (
	"encoding/json"
	"os"

	"github.com/gin-gonic/gin"
)

type users struct {
	ID            int     `json:"id"`
	Nome          string  `json:"nome"`
	Sobrenome     string  `json:"sobrenome"`
	Email         string  `json:"email"`
	Idade         string  `json:"idade"`
	Altura        float64 `json:"altura"`
	Ativo         bool    `json:"ativo"`
	DataDeCriacao string  `json:"dataDeCriacao"`
}

func main() {

	router := gin.Default()

	router.GET("/hello", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Olá, Katielly",
		})
	})

	router.GET("/users", getAll)

	router.Run()
}

func getAll(c *gin.Context) {
	data, err := os.ReadFile("users.json")

	if err != nil {
		panic("O arquivo indicado não foi encontrado ou está danificado ")
	}

	var u []users

	if err = json.Unmarshal([]byte(data), &u); err != nil {
		panic("Erro")
	}

	c.JSON(200, u)
}
