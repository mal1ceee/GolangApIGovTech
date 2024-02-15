package main

import (
	"GOLANGAPIGOVTECH/internal/config"
	"GOLANGAPIGOVTECH/internal/handler"
	"GOLANGAPIGOVTECH/internal/repository"
	"GOLANGAPIGOVTECH/internal/service"
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

func main() {
	cfg := config.LoadConfig()

	fmt.Println("Establishing connection to the database...")

	db, err := sqlx.Connect("postgres", cfg.DatabaseURL)
	if err != nil {
		fmt.Println("Failed to connect to the database.")
		panic(err)
	}
	fmt.Println("Connection to the database established.")

	repo := repository.NewRepository(db)
	svc := service.NewService(repo)
	h := handler.NewHandler(svc)

	router := gin.Default()
	h.RegisterRoutes(router)

	router.Run(cfg.ServerAddress)
}
