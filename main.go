package main

import (
	"anyWiki/common"
	"anyWiki/router"
	"log"
	"net/http"
)

func init() {
	//模板加载
	common.LoadTemplate()
}

func main() {
	server := http.Server{
		Addr: "127.0.0.1:8800",
	}
	//路由管理
	router.Router()

	if err := server.ListenAndServe(); err != nil {
		log.Println(err)
	}
}
