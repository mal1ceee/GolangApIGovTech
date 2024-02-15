package repository

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

type Student struct {
	ID     int    `db:"student_id"`
	Email  string `db:"email"`
	Status bool   `db:"status"`
	// Adjust fields according to your actual table structure
}

func ConnectToDB() {
	//Load configuration (e.g., database connection strings)
	// cfg, err := config.LoadConfig()
	// if err != nil {
	//     log.Fatal("Failed to load configuration: ", err)
	// }

	// Connect to the database
	db, err := sqlx.Connect("postgres", "postgres://postgres:password1@localhost:5432/postgres?sslmode=disable")
	if err != nil {
		log.Fatal("Failed to connect to the database: ", err)
	}

	log.Println("Querying the database...")

	if err := db.Ping(); err != nil {
		log.Fatalf("Unable to reach the database: %v", err)
	} else {
		log.Println("Successfully connected to the database.")
	}

	// Query the database for all student details, including the status
	var students []Student
	log.Println("About to query the database...")

	err = db.Select(&students, "SELECT student_id, email, status FROM students")
	if err != nil {
		log.Fatal("Failed to query the database: ", err)
	}
	log.Println("Query executed successfully.")

	log.Println("Fetched data from the database...")

	defer db.Close()

	// Log queried data
	for _, student := range students {
		// Adjust the log message to include the status
		log.Printf("ID: %d, Email: %s, Status: %t\n", student.ID, student.Email, student.Status)
	}

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
