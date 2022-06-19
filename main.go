package main

import (
	"github.com/GoLangDream/iceberg"
	_ "hot_news/initializers"
	"hot_news/works"
)

func main() {
	iceberg.InitApplication()
	works.Start()
	iceberg.StartApplication()
}
