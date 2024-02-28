package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	router := gin.Default()

	// make routes
	router.GET("", func(ctx *gin.Context) {
		ctx.String(http.StatusOK, "HELLO WORLD")
	})

	return router
}
