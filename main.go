package main

import (
	"fmt"
	"github.com/GoLangDream/iceberg"
	"github.com/GoLangDream/iceberg/environment"
	"github.com/GoLangDream/iceberg/log"
	_ "hot_news/initializers"
)

func main() {
	log.SetLevel(log.DebugLevel)
	iceberg.InitApplication()

	fmt.Printf("evn is %s\n", environment.Name())

	iceberg.StartApplication()
}
