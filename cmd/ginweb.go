package main

import (
	"fmt"
	"gin-web/pkg/setting"
	"gin-web/routers"
	"net/http"
)

func main() {
	router := routers.InitRouter()

	s := &http.Server{
		Addr:           fmt.Sprintf(":%d", setting.HttpPost),
		Handler:        router,
		ReadTimeout:    setting.ReadTimeOut,
		WriteTimeout:   setting.WriteTimeOut,
		MaxHeaderBytes: 1 << 20,
	}

	s.ListenAndServe()

}
