package views

import (
	"anyWiki/common"
	"anyWiki/config"
	"anyWiki/models"
	"net/http"
)

type IndexData struct {
	Title string `json:"title"`
	Desc  string `json:"desc"`
}

func (*HTMLApi) Index(w http.ResponseWriter, r *http.Request) {
	index := common.Template.Index

	var categories = []models.Category{
		{
			Cid:  1,
			Name: "go",
		},
	}

	var posts = []models.PostMore{
		{
			Pid:          1,
			Title:        "AnyWiki",
			Content:      "内容",
			UserName:     "tom",
			ViewCount:    123,
			CreateAt:     "2022-02-20",
			CategoryId:   1,
			CategoryName: "go",
			Type:         0,
		},
	}
	var hr = &models.HomeResponse{
		config.Cfg.Viewer,
		categories,
		posts,
		1,
		1,
		[]int{1},
		true,
	}

	index.WriteData(w, hr)

}
