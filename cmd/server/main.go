package main

import (
	// "fmt"
	"github.com/mal1ceee/GolangApIGovTech/internal/repository"
	// "net/http"
	// "GolangApIGovTech/internal/config"
	// "GolangApIGovTech/internal/handler"
	// "GolangApIGovTech/internal/service"

	// "github.com/gin-gonic/gin"
	// "github.com/jmoiron/sqlx"

	_ "github.com/lib/pq"
)

func main() {
	repository.ConnectToDB()
}
