package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/joho/godotenv"
)

func handleRequest(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" && r.URL.Path == "/health" {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		rsp := map[string]any{
			"status":    "ok",
			"timestamp": time.Now().Format(time.RFC3339),
		}
		json.NewEncoder(w).Encode(rsp)
	} else {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusNotFound)

		rsp := map[string]any{
			"error": "Not Found",
		}
		json.NewEncoder(w).Encode(rsp)
	}
}

func main() {
	// Load environment variables from .env file
	err := godotenv.Load("app.env")
	if err != nil {
		fmt.Println("Error loading app.env file")
	}

	http.HandleFunc("/", handleRequest)

	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}

	fmt.Printf("Server is running in %s mode on port %s\n", os.Getenv("ENVIRONMENT"), port)
	panic(http.ListenAndServe(":"+port, nil))
}
