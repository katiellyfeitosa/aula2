package main

import (
	"github.com/gin-gonic/gin"
	"github.com/katiellyfeitosa/bootcampGO/goWebII/parteII/exercicio/cmd/server/handler"
	"github.com/katiellyfeitosa/bootcampGO/goWebII/parteII/exercicio/internal/users"
)

func main() {
	repo := users.NewRepository()
	service := users.NewService(repo)
	u := handler.NewUser(service)

	r := gin.Default()
	groupRoute := r.Group("/users")
	groupRoute.POST("/", u.SaveUser())
	groupRoute.GET("/", u.GetAll())
	groupRoute.PUT("/:id", u.Update())
	groupRoute.DELETE("/:id", u.Delete())
	groupRoute.PATCH("/:id", u.UpdateLastNameAndAge())

	r.Run()
}
