// Exercício 2 - Validação de campo
// As validações dos campos devem ser implementadas no momento do envio do pedido, para
// isso devem ser seguidos os seguintes passos:
// 1. Todos os campos enviados na solicitação devem ser validados, todos os campos são
// obrigatórios
// 2. Caso algum campo não esteja completo, um código de erro 400 deve ser retornado
// com a mensagem “campo %s é obrigatório”.
// (Em %s deve ir o nome do campo que não está completo).
package main

import (
	"github.com/gin-gonic/gin"
)

type User struct {
	ID            int     `json:"id"`
	Nome          string  `json:"nome"`
	Sobrenome     string  `json:"sobrenome"`
	Email         string  `json:"email"`
	Idade         string  `json:"idade"`
	Altura        float64 `json:"altura"`
	Ativo         bool    `json:"ativo"`
	DataDeCriacao string  `json:"dataDeCriacao"`
}

var AllUsers []User
var lastID int

func main() {
	router := gin.Default()

	router.POST("/users", saveUser())
	router.GET("/users", getAllUsers())

	router.Run()
}

func saveUser() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.GetHeader("token")

		if token != "000000000" {
			c.JSON(401, gin.H{
				"error": "Token invalido",
			})
			return
		}

		var req User

		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(404, gin.H{
				"error": err.Error(),
			})
			return
		}

		lastID++
		req.ID = lastID
		AllUsers = append(AllUsers, req)

		c.JSON(200, gin.H{
			"data": req,
		})
	}
}

func getAllUsers() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(200, gin.H{
			"data": AllUsers,
		})
	}
}
