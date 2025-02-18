package controllers

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type RouteHandler struct {
	Route   string
	Method  string
	Handler gin.HandlerFunc
}

func RegisterHandler(handlers []RouteHandler, handler RouteHandler) []RouteHandler {
	handlers = append(handlers, handler)
	return handlers
}

func NewRouteHandler(method string, route string, handler gin.HandlerFunc) RouteHandler {
	return RouteHandler{
		Route:   route,
		Method:  method,
		Handler: handler,
	}
}

func New(r *gin.RouterGroup, handlers ...RouteHandler) {
	for _, closure := range handlers {
		h := closure
		switch h.Method {
		case http.MethodGet:
			r.GET(h.Route, h.Handler)
		case http.MethodPost:
			r.POST(h.Route, h.Handler)
		case http.MethodDelete:
			r.DELETE(h.Route, h.Handler)
		case http.MethodPatch:
			r.PATCH(h.Route, h.Handler)
		case http.MethodPut:
			r.PUT(h.Route, h.Handler)
		default:
			log.Fatal("Unimplemented")
		}
	}
}
