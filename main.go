/* making this file exec */
package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/krishpranav/restapi/user"
)

var usertable = []user.UserStruct{
	{User: "UserOne", Password: "PasswordTwo"},
	{User: "UserThree", Password: "PasswordThree"},
}

func main() {
	router := gin.Default()

	router.GET("/user", getUser)
	router.POST("/createuser", postUser)

	router.Run()
}

func getUser(c *gin.Context) {
	/* calling the var usertable and making giving response as a json with code 200 */
	c.IndentedJSON(http.StatusOK, usertable)
}

/* create new user using POST function */

func postUser(c *gin.Context) {
	var newUser user.UserStruct

	if err := c.BindJSON(&newUser); err != nil {
		return
	}

	usertable = append(usertable, newUser)
	c.IndentedJSON(http.StatusOK, newUser)
}

/**
curl http://localhost:8080/createuser \
    --include \
    --header "Content-Type: application/json" \
    --request "POST" \
    --data '{"user": "user", "password": "password"}'

use this command to create new user */
