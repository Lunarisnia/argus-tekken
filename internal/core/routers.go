package core

import (
	"context"
	"net/http"

	"github.com/Lunarisnia/argus-tekken/database/repo"
	"github.com/Lunarisnia/argus-tekken/internal/cheaters"
	"github.com/gin-gonic/gin"
)

func SetupRoute(ctx context.Context, r *gin.Engine, db *repo.Queries) {
	v1 := r.Group("/v1")

	// Setup Services
	cheaterService := cheaters.NewCheaterService(db)

	// Setup Routes
	cheaters.NewCheaterController(v1, cheaterService)

	v1.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})
}
