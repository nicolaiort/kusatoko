package main

import (
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Println("Error loading .env file")
		os.Exit(1)
	}

	port := os.Getenv("PORT")
	if port == "" {
		log.Println("PORT not found in env - defaulting to 8080")
		port = "8080"
	}

	rootMessage := os.Getenv("ROOT_MESSAGE")
	if rootMessage == "" {
		log.Println("ROOT_MESSAGE not found in env - defaulting to 'Hello, World!'")
		rootMessage = "Hello, World!"
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(rootMessage))
	})

	err = http.ListenAndServe(":"+port, nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
