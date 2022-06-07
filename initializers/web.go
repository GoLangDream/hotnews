package initializers

import (
	iceberg "github.com/GoLangDream/iceberg/web"
	"hot_news/web"
	_ "hot_news/web/controllers"
)

func init() {
	iceberg.RouterDraw = web.RouterDraw
}
