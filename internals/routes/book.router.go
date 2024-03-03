package routes

import (
	"bdpit/bookstore-go/internals/handlers"
	"bdpit/bookstore-go/internals/middlewares"
	"bdpit/bookstore-go/internals/repositories"

	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
)

func InitBookRouter(router *gin.Engine, db *sqlx.DB) {
	bookRouter := router.Group("/book")
	bookRepo := repositories.InitBookRepo(db)
	bookHandler := handlers.InitBookHandler(bookRepo)

	// localhost:800/book
	bookRouter.GET("", bookHandler.GetBooks)

	// localhost:8000/book/id
	bookRouter.GET("/:id", middlewares.CheckToken, bookHandler.GetBookById)

	// localhost:8000/book/new
	bookRouter.POST("/new",middlewares.CheckToken, bookHandler.CreateBook)

	// localhost:8000/book/id
	bookRouter.PATCH("/:id",middlewares.CheckToken, bookHandler.UpdateBook)

	//localhost:8000/book/id
	bookRouter.DELETE("/:id", middlewares.CheckToken, bookHandler.DeleteBookById)
}