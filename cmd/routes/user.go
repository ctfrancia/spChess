package routes

import (
	"net/http"

	"github.com/ctfrancia/spChess/cmd/handlers"
	"github.com/gin-gonic/gin"
)

func UserRoutes(g *gin.RouterGroup) {
	g.GET("/users", func(c *gin.Context) {
		c.JSON(
			http.StatusOK,
			gin.H{
				"code":    http.StatusOK,
				"message": "get all users endpoint, which I don't think that I'll need",
			},
		)
	})
	// g.GET("/users/:id", getUser)
	g.POST("/users", handlers.CreateUser)
	// g.PUT("/users/:id", updateUser)
	// g.DELETE("/users/:id", deleteUser)
}
