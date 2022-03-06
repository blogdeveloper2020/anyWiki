package router

import (
	"anyWiki/views"
	"net/http"
)

/**
返回数据分为三种类型
1.页面 viewers
2.api 数据(json)
3.静态资源
*/
func Router() {

	http.HandleFunc("/", views.HTML.Index)
	//静态资源加载
	http.Handle("/resource/", http.StripPrefix("/resource/", http.FileServer(http.Dir("public/resource/"))))
}
