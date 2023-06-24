package web

import (
	"github.com/GoLangDream/iceberg/web"
)

func RouterDraw(r *web.Router) {
	r.GET("/", "home#index")
	r.GET("/get_image_upload_url", "home#get_image_upload_url")
	r.GET("/image_url", "home#image_url")
	r.GET("/read_news", "home#read_news")
	r.GET("/update", "home#update")
	r.GET("/search_image", "home#search_image")
	r.GET("/web_content", "home#web_content")
	r.GET("/translate", "home#translate")
	r.GET("/feeds", "feeds#index")

	r.GET("/news", "news#index")
	r.GET("/news/read", "news#read")
}
