package mock

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func HandleEcho(ctx *gin.Context) {
	ctx.Data(http.StatusOK, "Content-Type: text/html; charset=utf-8", []byte("通了"))
}

func HandleGetUser(ctx *gin.Context) {
	ctx.AbortWithStatus(http.StatusInternalServerError)
}
