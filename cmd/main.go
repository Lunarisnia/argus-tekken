package main

import (
	"github.com/Lunarisnia/argus-tekken/internal/core"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	core.SetupRoute(r)
	r.Run(":3009") // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
