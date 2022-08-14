package works

import (
	"context"
	"github.com/GoLangDream/iceberg/log"
	"github.com/GoLangDream/iceberg/work"
	"github.com/mmcdole/gofeed"
	"hot_news/models"
	"hot_news/service"
	"time"
)

func syncRss(name, url string, needTranslate bool) {

	work.Register(name, "@hourly", func() {
		ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
		defer cancel()

		fp := gofeed.NewParser()

		feed, err := fp.ParseURLWithContext(url, ctx)
		if err != nil {
			log.Infof("获取 %s 的rss地址错误 %s", name, err)
			return
		}

		for _, item := range feed.Items {
			excerpt, image, _ := service.FetchWebContent(item.Link)

			if len(item.Link) > 1024 {
				log.Infof("文章 [%s] 的 url [%s] 超长", item.Title, item.Link)
				continue
			}

			news := &models.News{
				Title:       item.Title,
				Content:     excerpt,
				Url:         item.Link,
				SourceId:    item.GUID,
				SourceName:  name,
				Image:       image,
				IsTranslate: !needTranslate,
			}

			saveNews(news)
		}
	})
}
