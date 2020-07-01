package main

import (
	"fmt"
	"net/http"

	"github.com/PengWorkSpace/gin-example/pkg/setting"
	router "github.com/PengWorkSpace/gin-example/routers"
)

func main() {

	router := router.InitRouter()

	s := &http.Server{
		Addr:           fmt.Sprintf(":%d", setting.HTTPPort),
		Handler:        router,
		ReadTimeout:    setting.ReadTimeout,
		WriteTimeout:   setting.WriteTimeout,
		MaxHeaderBytes: 1 << 20,
	}
	s.ListenAndServe()
}
