package cheaters

import (
	"net/http"

	"github.com/Lunarisnia/argus-tekken/internal/controllers"
	"github.com/gin-gonic/gin"
)

var handlers []controllers.RouteHandler

func NewCheaterController(r *gin.RouterGroup) {
	cheater := r.Group("/cheater")

	ping()
	getAll()

	controllers.New(cheater, handlers...)
}

func getAll() {
	handler := func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"foo": "bar",
		})
	}

	handlers = controllers.RegisterHandler(handlers, controllers.RouteHandler{
		Route:   "/",
		Method:  http.MethodGet,
		Handler: handler,
	})
}

func ping() {
	handler := controllers.RouteHandler{
		Route:  "/ping",
		Method: http.MethodGet,
		Handler: func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"ping": "pong",
			})
		},
	}

	handlers = controllers.RegisterHandler(handlers, handler)
}
