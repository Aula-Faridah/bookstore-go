package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
)

func InitRouter(db *sqlx.DB) *gin.Engine {
	router := gin.Default()
	
	// make routes
	router.GET("", func(ctx *gin.Context) {
		ctx.String(http.StatusOK, "HELLO WORLD")
	})

	InitAuthRouter(router, db)
	InitBookRouter(router,db)
	
	return router
}
