package main

import (
	"github.com/GoLangDream/iceberg"
	"hot_news/initializers"
)

func main() {
	iceberg.InitApplication(&initializers.HotNewsApplicationConfig{})
	iceberg.StartApplication()
}
