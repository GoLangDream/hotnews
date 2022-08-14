package service

import (
	"fmt"
	"github.com/go-shiori/go-readability"
	"time"
)

func FetchWebContent(url string) (excerpt, image string, err error) {
	article, err := readability.FromURL(url, 10*time.Second)
	if err != nil {
		fmt.Printf("文章读取错误: %s\n url is [%s]\n", err, url)
		return "", "", err
	}
	return article.TextContent, article.Image, nil
}
