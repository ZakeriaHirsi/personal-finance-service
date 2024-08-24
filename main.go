package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"example.com/app"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	cognito "github.com/aws/aws-sdk-go/service/cognitoidentityprovider"
	"github.com/gofor-little/env"
)

func main() {
	log.SetOutput(os.Stdout)
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	conf := &aws.Config{Region: aws.String("eu-west-1")}
	sess, err := session.NewSession(conf)
	if err != nil {
		panic(err)
	}

	example := app.App{
		CognitoClient:   cognito.New(sess),
		UserPoolID:      goDotEnvVariable("COGNITO_USER_POOL_ID"),
		AppClientID:     goDotEnvVariable("COGNITO_APP_CLIENT_ID"),
		AppClientSecret: goDotEnvVariable("COGNITO_APP_CLIENT_SECRET"),
	}
	fmt.Println(goDotEnvVariable("COGNITO_USER_POOL_ID"))
	fmt.Println(goDotEnvVariable("COGNITO_APP_CLIENT_ID"))
	fmt.Println(goDotEnvVariable("COGNITO_APP_CLIENT_SECRET"))
	//Endpoint handlers
	http.HandleFunc("/health", HealthHandler)

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodPost:
			// Handle html form submission
			Call(&example, w, r)
		}
	})

	var port = goDotEnvVariable("PORT")
	if port == "" {
		port = "8080" // Default port
	}
	fmt.Println("Server listening on port", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}

func HelloHandler(w http.ResponseWriter, r *http.Request) {
	message := "Hello world!"
	encodeJSON(w, message)
}

func HealthHandler(w http.ResponseWriter, r *http.Request) {
	message := "Health check success from new ec2 instance"
	encodeJSON(w, message)
}

// Call routes POST requests
func Call(a *app.App, w http.ResponseWriter, r *http.Request) {

	switch r.URL.Path {
	case "/username":
		a.Username(w, r)
	default:
		http.Error(w, fmt.Sprintf("Handler for POST %s not found", r.URL.Path), http.StatusNotFound)
	}

	return
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
