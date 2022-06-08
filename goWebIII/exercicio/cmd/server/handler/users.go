package handler

import (
	"fmt"
	"os"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/katiellyfeitosa/bootcampGO/goWebIII/exercicio/internal/users"
)

type request struct {
	Nome      string  `json:"nome"`
	Sobrenome string  `json:"sobrenome"`
	Email     string  `json:"email"`
	Idade     string  `json:"idade"`
	Altura    float64 `json:"altura"`
}
type User struct {
	service users.Service
}

func NewUser(u users.Service) *User {
	return &User{
		service: u,
	}
}

func (c *User) GetAll() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		token := ctx.GetHeader("token")

		if token != os.Getenv("TOKEN") {
			ctx.JSON(401, gin.H{
				"error": "Token invalido",
			})
			return
		}

		u, err := c.service.GetAll()
		if err != nil {
			ctx.JSON(404, gin.H{
				"error": err.Error(),
			})
			return
		}
		ctx.JSON(200, u)
	}
}

func (c *User) SaveUser() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		token := ctx.GetHeader("token")

		if token != os.Getenv("TOKEN") {
			ctx.JSON(401, gin.H{
				"error": "Token invalido",
			})
			return
		}

		var req request
		if err := ctx.Bind(&req); err != nil {
			ctx.JSON(404, gin.H{
				"error": err.Error(),
			})
			return
		}

		u, err := c.service.SaveUser(req.Nome, req.Sobrenome, req.Email, req.Idade, req.Altura)
		if err != nil {
			ctx.JSON(404, gin.H{
				"error": err.Error(),
			})
			return
		}
		ctx.JSON(200, u)
	}
}

func (c *User) Update() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		token := ctx.GetHeader("token")

		if token != os.Getenv("TOKEN") {
			ctx.JSON(401, gin.H{
				"error": "Token invalido",
			})
			return
		}

		id, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
		if err != nil {
			ctx.JSON(400, gin.H{
				"error": err.Error(),
			})
			return
		}

		var req request
		if err := ctx.Bind(&req); err != nil {
			ctx.JSON(404, gin.H{
				"error": err.Error(),
			})
			return
		}

		if req.Nome == "" {
			ctx.JSON(400, gin.H{"error": "O nome é obrigatorio"})
			return
		}

		if req.Sobrenome == "" {
			ctx.JSON(400, gin.H{"error": "O sobrenome é obrigatorio"})
			return
		}

		if req.Email == "" {
			ctx.JSON(400, gin.H{"error": "O email é obrigatorio"})
			return
		}

		if req.Idade == "" {
			ctx.JSON(400, gin.H{"error": "A idade é obrigatorio"})
			return
		}

		if req.Altura == 0 {
			ctx.JSON(400, gin.H{"error": "A altura é obrigatorio"})
			return
		}

		u, err := c.service.Update(int(id), req.Nome, req.Sobrenome, req.Email, req.Idade, req.Altura)
		if err != nil {
			ctx.JSON(404, gin.H{
				"error": err.Error(),
			})
			return
		}
		ctx.JSON(200, u)
	}
}

func (c *User) UpdateLastNameAndAge() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		token := ctx.GetHeader("token")

		if token != os.Getenv("TOKEN") {
			ctx.JSON(401, gin.H{
				"error": "Token invalido",
			})
			return
		}

		id, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
		if err != nil {
			ctx.JSON(400, gin.H{
				"error": err.Error(),
			})
			return
		}

		var req request
		if err := ctx.Bind(&req); err != nil {
			ctx.JSON(404, gin.H{
				"error": err.Error(),
			})
			return
		}

		if req.Sobrenome == "" {
			ctx.JSON(400, gin.H{"error": "O sobrenome é obrigatorio"})
			return
		}

		if req.Idade == "" {
			ctx.JSON(400, gin.H{"error": "A idade é obrigatorio"})
			return
		}

		u, err := c.service.UpdateLastNameAndAge(int(id), req.Sobrenome, req.Idade)
		if err != nil {
			ctx.JSON(404, gin.H{
				"error": err.Error(),
			})
			return
		}
		ctx.JSON(200, u)
	}
}

func (c *User) Delete() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		token := ctx.GetHeader("token")

		if token != os.Getenv("TOKEN") {
			ctx.JSON(401, gin.H{
				"error": "Token invalido",
			})
			return
		}

		id, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
		if err != nil {
			ctx.JSON(400, gin.H{
				"error": err.Error(),
			})
			return
		}

		err = c.service.Delete(int(id))
		if err != nil {
			ctx.JSON(404, gin.H{
				"error": err.Error(),
			})
			return
		}
		ctx.JSON(200, gin.H{
			"data": fmt.Sprintf("O usuario %d foi removido", id),
		})
	}
}
