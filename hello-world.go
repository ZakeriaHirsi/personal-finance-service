package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gofor-little/env"
)

func main() {
	fmt.Println("hello world")

	var port = goDotEnvVariable("PORT")
	if port == "" {
		port = "8080" // Default port
	}
	log.SetOutput(os.Stdout)
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	http.HandleFunc("/", HelloHandler)
	http.HandleFunc("/health", HealthHandler)
	http.HandleFunc("/ap3", Ap3Handler)

	fmt.Println("Server listening on port", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}

func HelloHandler(w http.ResponseWriter, r *http.Request) {
	message := "Hello from GoLang API!"
	encodeJSON(w, message)
}

func HealthHandler(w http.ResponseWriter, r *http.Request) {
	message := "Health check success"
	encodeJSON(w, message)
}

func Ap3Handler(w http.ResponseWriter, r *http.Request) {
	message := "Hello Ap3s"
	encodeJSON(w, message)
}

func encodeJSON(w http.ResponseWriter, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	err := json.NewEncoder(w).Encode(data)
	if err != nil {
		log.Println("Error encoding data:", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
}

// use godot package to load/read the .env file and
// return the value of the key
func goDotEnvVariable(key string) string {

	// load .env file
	err := env.Load(".env")

	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	return os.Getenv(key)
}
