package main

import (
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"

	"github.com/joho/godotenv"
)

var rootMessage string
var port string

func handleRoot(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(rootMessage))
}

func handleHealthz(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("ok"))
}

func handleWhatsMyIP(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(strings.Split(r.RemoteAddr, ":")[0]))
}

func handleHeaders(w http.ResponseWriter, r *http.Request) {
	for k, v := range r.Header {
		w.Write([]byte(k + ": " + strings.Join(v, ", ") + "\n"))
	}
}
func handleStatus(w http.ResponseWriter, r *http.Request) {
	status, err := strconv.Atoi(r.URL.Path[len("/status/"):])
	if err != nil {
		http.Error(w, "Invalid status code", http.StatusBadRequest)
		return
	}
	http.Error(w, http.StatusText(status), status)
}

func initialize() {
	err := godotenv.Load()
	if err != nil && !os.IsNotExist(err) {
		log.Println("Error loading .env file:", err)
		os.Exit(1)
	}

	port = os.Getenv("PORT")
	if port == "" {
		log.Println("PORT not found in env - defaulting to 8080")
		port = "8080"
	}

	rootMessage = os.Getenv("ROOT_MESSAGE")
	if rootMessage == "" {
		log.Println("ROOT_MESSAGE not found in env - defaulting to 'Hello, World!'")
		rootMessage = "Hello, World!"
	}
}

func setupRoutes() *http.ServeMux {
	mux := http.NewServeMux()

	mux.HandleFunc("/", handleRoot)
	mux.HandleFunc("/healthz", handleHealthz)
	mux.HandleFunc("/whatsmyip", handleWhatsMyIP)
	mux.HandleFunc("/headers", handleHeaders)
	mux.HandleFunc("/status/", handleStatus)

	return mux
}

func main() {
	initialize()
	mux := setupRoutes()

	err := http.ListenAndServe(":"+port, mux)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
