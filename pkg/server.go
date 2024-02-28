package pkg

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func InitServer(router *gin.Engine) *http.Server {
	address := "localhost:8000"
	server := &http.Server{//use & bcs need its address
		Addr: address,
		Handler: router,
		WriteTimeout: time.Second * 10,
		IdleTimeout: time.Second * 10,
		ReadTimeout: time.Second * 10,

		} 
	return server
}