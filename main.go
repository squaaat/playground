package main

import (
	"fmt"
	"log"

	"github.com/squaaat/playground/internal/db"
	"github.com/squaaat/playground/internal/server"
	"github.com/squaaat/playground/internal/service/todo"
)


func main() {
	fmt.Println("wow")

	dbClient, err := db.New("0.0.0.0", "43306", "root", "pass", "mysql")
	if err != nil {
		log.Fatalf("1 %v", err)
	}
	results, err := dbClient.DB.Query("SELECT 1 + 1")
	if err != nil {
		panic(err.Error()) // proper error handling instead of panic in your app
	}
	fmt.Println(results)
	err = dbClient.InitDb()
	if err != nil {
		panic(err.Error()) // proper error handling instead of panic in your app

	}
	fmt.Println("success")

	server := server.New("0.0.0.0", "3000")
	todoService := todo.New(dbClient)
	todoService.RouteForHTTP(server)

	server.Listen()
}
