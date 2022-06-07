package web

import (
	"github.com/GoLangDream/iceberg/web"
)

func RouterDraw(r *web.Router) {
	r.GET("/", "home#index")
	r.GET("/update", "home#update")
	r.GET("/feeds", "feeds#index")
}
