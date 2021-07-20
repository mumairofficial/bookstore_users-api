package app

import "github.com/gin-gonic/gin"

var (
	router = gin.Default()
)

func StartApplication() {
	// start logic
	mapUrls()
	router.Run(":8080")
}