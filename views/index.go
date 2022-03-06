package views

import (
	"anyWiki/config"
	"anyWiki/models"
	"html/template"
	"log"
	"net/http"
	"time"
)

type IndexData struct {
	Title string `json:"title"`
	Desc  string `json:"desc"`
}

func IsODD(num int) bool {
	return num%2 == 0
}
func GetNextName(strs []string, index int) string {
	return strs[index+1]
}

func Date(layout string) string {
	return time.Now().Format(layout)
}

func (*HTMLApi) Index(w http.ResponseWriter, r *http.Request) {
	var indexData IndexData
	indexData.Title = "anyWiki"
	indexData.Desc = "你好,欢迎来到anyWiki"
	t := template.New("index.html")
	//1.get current path
	path := config.Cfg.System.CurrentDir
	home := path + "/template/home.html"
	header := path + "/template/layout/header.html"
	footer := path + "/template/layout/footer.html"
	personal := path + "/template/layout/personal.html"
	pagination := path + "/template/layout/pagination.html"
	postList := path + "/template/layout/post-list.html"
	t.Funcs(template.FuncMap{"isODD": IsODD, "getNextName": GetNextName, "date": Date})
	t, err := t.ParseFiles(path+"/template/index.html",
		home, header, footer, personal, pagination, postList)
	if err != nil {
		log.Println(err)
	}
	//2.output the indexData to html file
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
	t.Execute(w, hr)
}
