package mock

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func NewRouter() *gin.Engine {
	router := gin.New()

	g := router.Group("/")
	g.Handle(http.MethodGet, "/echo", HandleEcho)
	g.Handle(http.MethodGet, "/user", HandleGetUser)
	return router
}
