package web

import (
	. "github.com/GoLangDream/iceberg/web"
)

func RouterDraw(r *Router) {
	r.GET("/", "home#index")
}
