/* making this file exec */
package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type user struct {
	UserName string `json:"username"`
	Password string `json:"password"`
}

var usertable = []user{
	{UserName: "UserOne", Password: "StrongPassword"},
	{UserName: "UserTwo", Password: "StrongPasswordTwo"},
}

func main() {
	router := gin.Default()

	router.GET("/user", getUser)

	router.Run()
}

func getUser(c *gin.Context) {
	/* calling the var usertable and making giving response as a json with code 200 */
	c.IndentedJSON(http.StatusOK, usertable)
}
