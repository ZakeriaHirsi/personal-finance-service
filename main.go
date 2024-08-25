package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/cognitoidentityprovider"
	"github.com/aws/aws-sdk-go-v2/service/cognitoidentityprovider/types"
	"github.com/gofor-little/env"
)

func main() {
	var port = goDotEnvVariable("PORT")
	if port == "" {
		port = "8080" // Default port
	}
	log.SetOutput(os.Stdout)
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	//Handlers
	http.HandleFunc("/users", UserPoolHandler)

	fmt.Println("Server listening on port", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}

func UserPoolHandler(w http.ResponseWriter, r *http.Request) {
	sdkConfig, err := config.LoadDefaultConfig(context.TODO())
	if err != nil {
		fmt.Println("Couldn't load default configuration. Have you set up your AWS account?")
		fmt.Println(err)
		return
	}
	cognitoClient := cognitoidentityprovider.NewFromConfig(sdkConfig)

	var users []types.UserType
	var userPoolId = goDotEnvVariable("COGNITO_USER_POOL_ID")
	usersPaginator := cognitoidentityprovider.NewListUsersPaginator(cognitoClient, &cognitoidentityprovider.ListUsersInput{UserPoolId: aws.String(userPoolId)})
	for usersPaginator.HasMorePages() {
		output, err := usersPaginator.NextPage(context.TODO())
		if err != nil {
			log.Printf("Couldn't get users. Here's why: %v\n", err)
		} else {
			users = append(users, output.Users...)
		}
	}

	if len(users) == 0 {
		fmt.Println("You don't have any user pools!")
	} else {
		for _, user := range users {
			fmt.Printf("\t%v\n", *user.Username)
		}
	}

	encodeJSON(w, users)
}

func goDotEnvVariable(key string) string {
	// load .env file
	err := env.Load(".env")
	if err != nil {
		log.Fatalf("Error loading .env file")
	}
	return os.Getenv(key)
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
