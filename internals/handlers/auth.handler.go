package handlers

import (
	"bdpit/bookstore-go/internals/models"
	"bdpit/bookstore-go/internals/repositories"
	"bdpit/bookstore-go/pkg"
	_ "hash"
	"log"
	"net/http"

	"github.com/alexedwards/argon2id"
	"github.com/gin-gonic/gin"
	_ "golang.org/x/net/context/ctxhttp"
)

type AuthHandler struct {
	*repositories.AuthRepo
}

func InitAuthHandler(a *repositories.AuthRepo) *AuthHandler {
	return &AuthHandler{a}
}

func (a *AuthHandler) Register(ctx *gin.Context) {
	// ambil body
	body := models.AuthModel{}
	if err := ctx.ShouldBind(&body); err != nil {
		log.Println(err.Error())
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	result, err := a.FindByEmail(body)
	if err != nil {
		log.Println(err.Error())
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	// checking duplicate email
	if len(result) > 0 {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "Email is already regitered!",
		})
		return
	}

	hash, err := argon2id.CreateHash(body.Password, argon2id.DefaultParams)
	if err != nil {
		log.Println(err.Error())
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	if err := a.SaveUser(models.AuthModel{
		Email:    body.Email,
		Password: hash,
	}); err != nil {
		log.Println(err.Error())
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{
		"message": "Success Registered!",
	})
}

func (a *AuthHandler) Login(ctx *gin.Context) {
	body := models.AuthModel{}
	if err := ctx.ShouldBind(&body); err != nil {
		log.Println(err.Error())
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	result, err := a.FindByEmail(body)
	if err != nil {
		log.Println(err.Error())
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	if len(result) == 0 {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "Email isn't registered!",
		})
		return
	}

	match, err := argon2id.ComparePasswordAndHash(body.Password, result[0].Password)
	if err != nil {
		log.Println(err.Error())
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	if !match {
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"message": "Bad Credentials!",
		})
		return
	}

	payload := pkg.NewPayload(body.Email)
	token, err := payload.CreateToken()
	if err != nil {
		log.Println(err.Error())
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message":"Login Success!",
		"token": token,
	})
}
