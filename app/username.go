package app

import (
	"fmt"
	"net/http"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/awserr"

	cognito "github.com/aws/aws-sdk-go/service/cognitoidentityprovider"
)

func (a *App) Username(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	username := r.Form.Get("username")
	fmt.Println(username)
	fmt.Println(a.UserPoolID)
	fmt.Println(a.AppClientID)
	fmt.Println(a.AppClientSecret)
	fmt.Println(a.CognitoClient.ClientInfo.Endpoint)
	_, err := a.CognitoClient.AdminGetUser(&cognito.AdminGetUserInput{
		UserPoolId: aws.String(a.UserPoolID),
		Username:   aws.String(username),
	})
	fmt.Println("cognito client done work")
	fmt.Println(err.Error())
	if err != nil {
		awsErr, ok := err.(awserr.Error)
		if ok {
			if awsErr.Code() == cognito.ErrCodeUserNotFoundException {
				m := fmt.Sprintf("Username %s is free!", username)
				fmt.Println("username is free")
				http.Redirect(w, r, fmt.Sprintf("/username?message=%s", m), http.StatusSeeOther)
				return
			}
		} else {
			http.Redirect(w, r, "Something went wrong", http.StatusSeeOther)
			return
		}
	}

	m := fmt.Sprintf("Username %s is taken.", username)
	fmt.Println("username is taken")
	http.Redirect(w, r, fmt.Sprintf("/username?message=%s", m), http.StatusSeeOther)
}
