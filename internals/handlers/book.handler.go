package handlers

import (
	"bdpit/bookstore-go/internals/models"
	"bdpit/bookstore-go/internals/repositories"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type BookHandler struct {
	*repositories.BookRepo
}

func InitBookHandler(b *repositories.BookRepo) *BookHandler {
	return &BookHandler{b}
}

func (b *BookHandler) GetBooks(ctx *gin.Context) {
	result, err := b.FindAll()
	if err != nil {
		log.Println(err.Error())
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"message": "Success Get Book",
		"data":    result,
	})
}

func (b *BookHandler) GetBookById(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error":"Invalid book ID.",
		})
		return
	}
	
	result, err := b.FindById(id)

	if err != nil {
		log.Println(err.Error())
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to get book.",
		})
		return
	}

	if len(result) == 0 {
		ctx.JSON(http.StatusNotFound, gin.H{
			"error" : "Book not found.",
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"messages":"Success get book!",
		"data":result[0],
	})
}

func (b *BookHandler) CreateBook(ctx *gin.Context) {
	body := models.BookModel{}
	if err := ctx.ShouldBind(&body); err != nil {
		log.Println(err.Error())
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	if err := b.SaveBook(body); err != nil {
		log.Println(err.Error())
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	result, err := b.FindAll()
	if err != nil {
		log.Println(err.Error())
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{
		"message": "Success Save Book",
		"data":    result,
	})
}

func (b *BookHandler) UpdateBook(ctx *gin.Context) {
	body := models.BookModel{}
	// var body models.BookModel
	if err := ctx.ShouldBind(&body); err != nil {
		log.Println(err.Error())
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid request body",
		})
		return
	}

	id, err := strconv.Atoi(ctx.Param("id"))

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid book ID",
		})
		return
	}

	body.Id = id

	if err := b.BookRepo.UpdateBook(body); err != nil {
		log.Println(err.Error())
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to update book",
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "Success Update Book",
		"data":    body,
	})
}

func (b *BookHandler) DeleteBookById(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error":"Invalid book ID.",
		})
		return
	}

	if err := b.BookRepo.DeleteBookById(id); err != nil {
		log.Println(err.Error())
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error":"Failed to delete book.",
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message":"Book deleted success!",
	})
	
}