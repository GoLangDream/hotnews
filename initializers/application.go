package initializers

import (
	iceberg "github.com/GoLangDream/iceberg/web"
	"hot_news/web"
	"os"
)

type HotNewsApplicationConfig struct {
}

func (app *HotNewsApplicationConfig) HomePath() string {
	path, _ := os.Getwd()
	return path
}

func (app *HotNewsApplicationConfig) RouterDraw() func(r *iceberg.Router) {
	return web.RouterDraw
}
