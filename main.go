package main

import (
	"github.com/GoLangDream/iceberg"
	_ "hot_news/initializers"
)

func main() {
	iceberg.InitApplication()
	iceberg.StartApplication()
}
