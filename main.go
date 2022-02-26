package main

import (
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
	t, _ = t.ParseFiles(path + "/template/index.html")
	//2.output the indexData to html file
	t.Execute(w, indexData)
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
