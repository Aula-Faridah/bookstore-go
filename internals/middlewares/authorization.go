package middlewares

import (
	"bdpit/bookstore-go/pkg"
	_"go/token"
	"log"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func CheckToken(ctx *gin.Context) {
	// ambil header authorization
	bearerToken := ctx.GetHeader("Authorization")
	// Bearer token
	if bearerToken == "" {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"message":"Please Login First",
		})
		return 
	}

	if !strings.Contains(bearerToken, "Bearer "){
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"message":"Invalid Authorization",
		})
		return 
	}

	token := strings.Replace(bearerToken, "Bearer ", "", -1)
	_, err := pkg.VerifyToken(token)
	if err != nil {
		if strings.Contains(err.Error(), "Expired") {
			log.Println(err.Error())
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"message": "Please Login Again!",
			})
			return 
		}
		log.Println(err.Error())
			ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
				"error": err.Error(),
			})
			return
	}
	ctx.Next()
}