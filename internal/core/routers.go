package core

import (
	"net/http"

	"github.com/Lunarisnia/argus-tekken/internal/cheaters"
	"github.com/gin-gonic/gin"
)

func SetupRoute(r *gin.Engine) {
	v1 := r.Group("/v1")

	// Setup Routes
	cheaters.NewCheaterController(v1)

	v1.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})
}
