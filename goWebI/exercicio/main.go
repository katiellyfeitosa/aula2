// Exercício 2 - Get one endpoint

// Gere um novo endpoint que nos permita buscar um único resultado do array de temas.
// Usando parâmetros de caminho o endpoint deve ser /theme/:id (lembre-se que o tema
// sempre tem que ser plural). Uma vez que o id é recebido, ele retorna a posição
// correspondente.
// 1. Gere uma nova rota.
// 2. Gera um manipulador para a rota criada.
// 3. Dentro do manipulador, procure o item que você precisa.
// 4. Retorna o item de acordo com o id.
// Se você não encontrou nenhum elemento com esse id retorne como código de resposta 404.
package main

import (
	"encoding/json"
	"net/http"
	"os"
	"strconv"

	"github.com/gin-gonic/gin"
)

type User struct {
	ID            int    `json:"id"`
	Nome          string `json:"nome"`
	Sobrenome     string `json:"sobrenome"`
	Email         string `json:"email"`
	Idade         string `json:"idade"`
	Altura        string `json:"altura"`
	Ativo         string `json:"ativo"`
	DataDeCriacao string `json:"dataDeCriacao"`
}

func main() {
	router := gin.Default()

	router.GET("/users", getAll)
	router.GET("/users/:id", getBy)

	router.Run()
}

func getAll(c *gin.Context) {
	data, err := os.ReadFile("users.json")

	if err != nil {
		panic("O arquivo indicado não foi encontrado ou está danificado ")
	}

	var allUsers []User

	if err = json.Unmarshal(data, &allUsers); err != nil {
		panic(err)
	}

	filteredUser := filter(c, allUsers)

	if len(filteredUser) != 0 {
		c.JSON(http.StatusOK, gin.H{
			"data": filteredUser,
		})
		return
	} else {
		c.JSON(http.StatusOK, allUsers)
	}
}

func filter(c *gin.Context, allUsers []User) []User {
	var filteredUser []User

	id := c.Query("id")
	nome := c.Query("nome")
	sobrenome := c.Query("sobrenome")
	email := c.Query("email")
	idade := c.Query("idade")
	altura := c.Query("altura")

	parsedID, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Erro interno, tente novamente",
		})
		return filteredUser
	}

	var appendUser bool

	for _, user := range allUsers {
		if id != "" {
			if user.ID == parsedID {
				appendUser = true
			}
		}
		if nome != "" {
			if user.Nome == nome {
				appendUser = true
			}
		}
		if sobrenome != "" {
			if user.Sobrenome == sobrenome {
				appendUser = true
			}
		}
		if email != "" {
			if user.Email == email {
				appendUser = true
			}
		}
		if idade != "" {
			if user.Idade == idade {
				appendUser = true
			}
		}
		if altura != "" {
			if user.Altura == altura {
				appendUser = true
			}
		}

		if appendUser == true {
			filteredUser = append(filteredUser, user)
		}
	}
	return filteredUser
}

func getBy(c *gin.Context) {
	data, err := os.ReadFile("users.json")

	if err != nil {
		panic("O arquivo indicado não foi encontrado ou está danificado ")
	}

	id := c.Param("id")
	parsedID, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Erro interno, tente novamente",
		})
		return
	}

	var allUsers []User

	if err = json.Unmarshal(data, &allUsers); err != nil {
		panic(err)
	}

	var filteredUser []User

	for _, user := range allUsers {
		if id != "" {
			if user.ID == parsedID {
				filteredUser = append(filteredUser, user)
			}
		}
	}

	if len(filteredUser) != 0 {
		c.JSON(http.StatusOK, gin.H{
			"data": filteredUser,
		})
		return
	} else {
		c.JSON(404, "Usuario não encontrado!")
	}
}
