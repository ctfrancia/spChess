package handlers

import (
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/gin-gonic/gin"
)

// CreateUser is the handler function for creating a new user
func CreateUser(c *gin.Context) {
	jsonData, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		fmt.Println(err)
	}
	c.JSON(
		http.StatusOK,
		gin.H{
			"code":    http.StatusOK,
			"message": "inside CreateUser",
		},
	)
	fmt.Printf("incomming body: %v", string(jsonData))
}
