package main

import (
	"fmt"
	"net/http"

	"github.com/PengWorkSpace/gin-example/pkg/setting"
	"github.com/gin-gonic/gin"
)

func main() {

	router := gin.Default()
	router.GET("/test", func(c *gin.Context) {

		c.JSON(200, gin.H{
			"message": "this is a test message",
		})
	})

	s := &http.Server{
		Addr:           fmt.Sprintf(":%d", setting.HTTPPort),
		Handler:        router,
		ReadTimeout:    setting.ReadTimeout,
		WriteTimeout:   setting.WriteTimeout,
		MaxHeaderBytes: 1 << 20,
	}
	s.ListenAndServe()
}
