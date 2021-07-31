package main

import (
	"fmt"
	"log"

	"github.com/squaaat/playground/internal/db"
)

func main() {
	fmt.Println("wow")

	client, err := db.New("0.0.0.0", "43306", "root", "pass", "mysql")
	if err != nil {
		log.Fatalf("1 %v", err)
	}
	results, err := client.DB.Query("SELECT 1 + 1")
	if err != nil {
		panic(err.Error()) // proper error handling instead of panic in your app
	}
	fmt.Println(results)
	err = client.InitDb()
	if err != nil {
		panic(err.Error()) // proper error handling instead of panic in your app

	}
	fmt.Println("success")

	err = client.Create("1")
	err = client.Create("1")
	err = client.Create("1")
	err = client.Create("1")
	err = client.Create("1")
	err = client.Create("1")
	err = client.Create("1")

	if err != nil {
		fmt.Println(err)
	}

	res, err := client.GetAll()

	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(res)

}
