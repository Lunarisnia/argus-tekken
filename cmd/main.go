package main

import (
	"context"

	"github.com/Lunarisnia/argus-tekken/internal/core"
	"github.com/Lunarisnia/argus-tekken/internal/db"
	"github.com/gin-gonic/gin"
)

// TODO: Get Cheater API
// TODO: Get All cheaters

// TODO: Get Evidence based on Polaris ID
// TODO: Get all evidences
func main() {
	ctx := context.Background()

	r := gin.Default(func(e *gin.Engine) {
	})
	q, conn := db.ConnectDatabase(ctx)
	defer conn.Close(ctx)

	core.SetupRoute(ctx, r, q)
	r.Run(":3009") // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
