package main

import (
   
    "log"

    "github.com/jmoiron/sqlx"
    _ "github.com/lib/pq"
)

func main() {
    
    db, err := sqlx.Connect("postgres", "postgres://postgres:TOUCHthesky123!@localhost:5432/postgres?sslmode=disable")
    if err != nil {
        log.Fatal("Failed to connect to the database: ", err)
    }
    defer db.Close()

    // Ping the database to ensure it's accessible
    if err := db.Ping(); err != nil {
        log.Fatalf("Unable to reach the database: %v", err)
    } else {
        log.Println("Successfully connected to the database.")
    }

    
}