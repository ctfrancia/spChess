package routes

func UserRoutes(g *gin.RouterGroup) {
	g.GET("/users", getUsers)
	// g.GET("/users/:id", getUser)
	// g.POST("/users", createUser)
	// g.PUT("/users/:id", updateUser)
	// g.DELETE("/users/:id", deleteUser)
}
