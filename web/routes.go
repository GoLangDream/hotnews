package web

import (
	"github.com/GoLangDream/iceberg/web"
)

func RouterDraw(r *web.Router) {
	r.GET("/", "home#index")
	r.GET("/news", "home#news")
	r.GET("/update", "home#update")
	r.GET("/web_content", "home#web_content")
	r.GET("/translate", "home#translate")
	r.GET("/feeds", "feeds#index")
}
