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
	ctl.newCheater()
	ctl.newEvidence()

	controllers.New(cheater, handlers...)
}

type cheaterCtl struct {
	cs CheaterService
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

func (ch *cheaterCtl) newEvidence() {
	h := controllers.RouteHandler{
		Method: http.MethodPost,
		Route:  "/evidence",
		Handler: func(c *gin.Context) {
			newEvidence := cheaterparams.NewEvidence{}
			if err := c.ShouldBindBodyWithJSON(&newEvidence); err != nil {
				log.Println(err)
				c.JSON(http.StatusBadRequest, gin.H{
					"message": "bad request",
				})
				return
			}

			err := ch.cs.NewEvidence(c.Request.Context(), newEvidence)
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
