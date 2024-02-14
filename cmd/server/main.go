package main

import (
    // "fmt"
    "log"
    // "net/http"
    // "GolangApIGovTech/internal/config"
    // "GolangApIGovTech/internal/handler"
    // "GolangApIGovTech/internal/repository"
    // "GolangApIGovTech/internal/service"

    // "github.com/gin-gonic/gin"
    "github.com/jmoiron/sqlx"
    _ "github.com/lib/pq"
)

func main() {
    // Load configuration (e.g., database connection strings)
    // cfg, err := config.LoadConfig()
    // if err != nil {
    //     log.Fatal("Failed to load configuration: ", err)
    // }

    // Connect to the database
    db, err := sqlx.Connect("postgres", "postgres://postgres:TOUCHthesky123!@localhost:5432/postgres?sslmode=disable")
    if err != nil {
        log.Fatal("Failed to connect to the database: ", err)
    }
    defer db.Close()

//     // Initialize repository
//     repo := repository.NewRepository(db)

//     // Initialize service
//     svc := service.NewService(repo)

//     // Initialize HTTP handler
//     h := handler.NewHandler(svc)

//     // Initialize Gin router
//     router := gin.Default()

//     // Set up routes
//     setupRoutes(router, h)

//     // Start the server
//     if err := router.Run(":8080"); err != nil {
//         log.Fatal("Failed to run the server: ", err)
//     }
// }

// func setupRoutes(router *gin.Engine, h *handler.Handler) {
//     // Define your routes here
//     router.GET("/example", h.ExampleHandler)
 }