package main

import (
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/katiellyfeitosa/bootcampGO/goWebIII/exercicio/cmd/server/handler"
	"github.com/katiellyfeitosa/bootcampGO/goWebIII/exercicio/internal/users"
	"github.com/katiellyfeitosa/bootcampGO/goWebIII/exercicio/pkg/store"
)

func main() {
	_ = godotenv.Load("../../.env")
	db := store.New(store.FileType, "users.json")
	repo := users.NewRepository(db)
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
