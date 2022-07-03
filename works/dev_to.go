package works

import (
	"github.com/GoLangDream/iceberg/log"
	"github.com/GoLangDream/iceberg/work"
	"github.com/mmcdole/gofeed"
	"hot_news/models"
	"hot_news/service"
)

var fp = gofeed.NewParser()

func init() {
	work.Register("dev.to", "@hourly", syncDevTo)
}

func syncDevTo() {
	feed, err := fp.ParseURL("https://dev.to/feed")
	if err != nil {
		log.Infof("获取 dev.to 的rss地址错误 %s", err)
	}

	for _, item := range feed.Items {
		if checkNewsExists("dev.to", item.GUID) {
			continue
		}

		cnTitle := service.TranslateString(item.Title)

		news := models.News{
			Title:      item.Title,
			CnTitle:    cnTitle,
			Content:    "",
			Url:        item.Link,
			SourceId:   item.GUID,
			SourceName: "dev.to",
		}

		saveNews(news)
	}
}
