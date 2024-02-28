package routes

import (
	"bdpit/bookstore-go/internals/handlers"
	"bdpit/bookstore-go/internals/repositories"

	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
)

func InitAuthRouter(router *gin.Engine, db *sqlx.DB) {
	// Bikin Subrouter
	authRouter := router.Group("/auth")

	authRepo := repositories.InitAuthRepo(db)
	authHandler := handlers.InitAuthHandler(authRepo)

	// Bikin rute
	authRouter.POST("/new", authHandler.Register)
	authRouter.POST("", authHandler.Login)
}