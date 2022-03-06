package main

import (
	"anyWiki/config"
	"anyWiki/models"
	"html/template"
	"log"
	"net/http"
	"os"
)

type IndexData struct {
	Title string `json:"title"`
	Desc  string `json:"desc"`
}

func index(w http.ResponseWriter, r *http.Request) {
	var indexData IndexData
	indexData.Title = "anyWiki"
	indexData.Desc = "你好,欢迎来到anyWiki"
	t := template.New("index.html")
	//1.get current path
	path, _ := os.Getwd()
	home := path + "/template/home.html"
	header := path + "/template/layout/header.html"
	footer := path + "/template/layout/footer.html"
	personal := path + "/template/layout/personal.html"
	pagination := path + "/template/layout/pagination.html"
	poster := path + "/template/layout/poster.html"

	t, _ = t.ParseFiles(path+"/template/index.html",
		home, header, footer, personal, pagination, poster)
	//2.output the indexData to html file
	viewer:=config.Viewer{
		Title: "AnyWiki",
		Description: "AnyWiki"
		Logo: "",
	}
	var hr = &models.HomeResponse{}
	t.Execute(w, hr)
}

func main() {
	server := http.Server{
		Addr: "127.0.0.1:8800",
	}
	http.HandleFunc("/", index)

	if err := server.ListenAndServe(); err != nil {
		log.Println(err)
	}
}
