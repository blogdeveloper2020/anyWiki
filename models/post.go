package models

import (
	"anyWiki/config"
	"html/template"
	"time"
)

type Post struct {
	Pid        int       `json:"pid"`
	Title      string    `json:"title"`
	Slug       string    `json:"slug"`
	Content    string    `json:"content"`
	Markdown   string    `json:"markdown"`
	CategoryId int       `json:"categoryId"`
	UserId     int       `json:"userId"`
	ViewCount  int       `json:"viewCount"`
	Type       int       `json:"type"`
	CreateAt   time.Time `json:"createAt"`
	Update     time.Time `json:"update"`
}

type PostMore struct {
	Pid          int           `json:"pid"`
	Title        string        `json:"title"`
	Slug         string        `json:"slug"`
	Content      template.HTML `json:"content"`
	CategoryId   int           `json:"categoryId"`
	CategoryName string        `json:"categoryName"`
	UserId       int           `json:"userId"`
	UserName     string        `json:"userName"`
	ViewCount    int           `json:"viewCount"`
	Type         int           `json:"type"`
	CreateAt     string        `json:"createAt"`
	UpdateAt     string        `json:"updateAt"`
}

type PostReq struct {
	Pid      int    `json:"pid"`
	Title    string `json:"title"`
	Slug     string `json:"slug"`
	Content  string `json:"content"`
	Markdown string `json:"markdown"`
	Category int    `json:"category"`
	UserId   int    `json:"userId"`
	Type     int    `json:"type"`
}
type SearchResp struct {
	Pid   int    `json:"pid"`
	Title string `json:"title"`
}
type PostRes struct {
	config.Viewer
	config.SystemConfig
	Article PostMore
}
