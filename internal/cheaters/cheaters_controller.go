package cheaters

import (
	"log"
	"net/http"

	"github.com/Lunarisnia/argus-tekken/internal/cheaters/cheaterparams"
	"github.com/Lunarisnia/argus-tekken/internal/controllers"
	"github.com/gin-gonic/gin"
)

var handlers []controllers.RouteHandler

func NewCheaterController(r *gin.RouterGroup, cs CheaterService) {
	cheater := r.Group("/cheater")

	ctl := cheaterCtl{
		cs: cs,
	}
	ctl.ping()
	ctl.getAll()
	ctl.newCheater()

	controllers.New(cheater, handlers...)
}

type cheaterCtl struct {
	cs CheaterService
}

func (ch *cheaterCtl) getAll() {
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

func (ch *cheaterCtl) ping() {
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

func (ch *cheaterCtl) newCheater() {
	h := controllers.RouteHandler{
		Method: http.MethodPost,
		Route:  "/",
		Handler: func(c *gin.Context) {
			newCheater := cheaterparams.NewCheater{}
			if err := c.ShouldBindBodyWithJSON(&newCheater); err != nil {
				log.Println(err)
				c.JSON(http.StatusBadRequest, gin.H{
					"message": "bad request",
				})
				return
			}

			err := ch.cs.NewCheater(c.Request.Context(), newCheater)
			if err != nil {
				log.Println(err)
				c.JSON(http.StatusInternalServerError, gin.H{
					"message": "Failed to insert",
				})
				return
			}

			c.JSON(http.StatusOK, gin.H{
				"message": "OK",
			})
		},
	}

	handlers = controllers.RegisterHandler(handlers, h)
}
