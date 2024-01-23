package router

import "github.com/gin-gonic/gin"

func Initialize() {
	// Initialize router
	r := gin.Default()

	// Initialize routes
	initializeRoutes(r)
	// run the server
	r.Run() // listen and serve on 0.0.0.0:8080

}
