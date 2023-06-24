package main

import (
	"fmt"
	"github.com/GoLangDream/iceberg"
	"github.com/GoLangDream/iceberg/environment"
	"github.com/GoLangDream/iceberg/log"
	_ "hot_news/initializers"
	"os"
	"strconv"
)

func main() {
	pid := os.Getpid()
	err := os.WriteFile("application.pid", []byte(strconv.Itoa(pid)), 0644)
	if err != nil {
		fmt.Println("application.pid 文件写入错误 ")
		return
	}
	log.SetLevel(log.DebugLevel)
	iceberg.InitApplication()

	fmt.Printf("evn is %s\n", environment.Name())

	iceberg.StartApplication()
}
