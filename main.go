package main

import (
	"encoding/json"
	"log"
	"net/http"
)

type IndexData struct {
	Title string `json:"title"`
	Desc  string `json:"desc"`
}

func index(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var indexData IndexData
	indexData.Title = "anyWiki"
	indexData.Desc = "a wiki web application"
	jsonStr, _ := json.Marshal(indexData)
	w.Write(jsonStr)
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
