package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/foyez/go/api"
	"github.com/joho/godotenv"
)

func main() {
	// Load environment variables from .env file
	err := godotenv.Load("app.env")
	if err != nil {
		fmt.Println("Error loading app.env file")
	}

	http.HandleFunc("/", api.HandleRequest)

	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}

	fmt.Printf("Server is running in %s mode on port %s\n", os.Getenv("ENVIRONMENT"), port)
	panic(http.ListenAndServe(":"+port, nil))
}
