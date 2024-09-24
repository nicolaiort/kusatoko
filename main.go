package main

import (
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil && !os.IsNotExist(err) {
		log.Println("Error loading .env file:", err)
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

	http.HandleFunc("/healthz", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("ok"))
	})

	http.HandleFunc("/whatsmyip", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(strings.Split(r.RemoteAddr, ":")[0]))
	})

	http.HandleFunc("/headers", func(w http.ResponseWriter, r *http.Request) {
		for k, v := range r.Header {
			w.Write([]byte(k + ": " + strings.Join(v, ", ") + "\n"))
		}
	})

	http.HandleFunc("/status/", func(w http.ResponseWriter, r *http.Request) {
		status, err := strconv.Atoi(r.URL.Path[len("/status/"):])
		if err != nil {
			http.Error(w, "Invalid status code", http.StatusBadRequest)
			return
		}
		http.Error(w, http.StatusText(status), status)
	})

	err = http.ListenAndServe(":"+port, nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
