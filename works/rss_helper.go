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

func syncRss(name, url string, needTranslate ...bool) {

	_needTranslate := true

	if len(needTranslate) > 0 && needTranslate[0] == false {
		_needTranslate = false
	}

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
			if checkNewsExists(name, item.GUID) {
				continue
			}

			cnTitle := item.Title
			cnDescription := ""
			excerpt, image := "", ""

			if _needTranslate {
				cnTitle = service.TranslateString(item.Title)
				excerpt, image, _ = service.FetchWebContent(item.Link)
				if excerpt != "" {
					cnDescription = service.TranslateString(excerpt)
				}
			}

			if len(item.Link) > 250 {
				log.Infof("文章 [%s] 的 url [%s] 超长", cnTitle, item.Link)
				continue
			}

			news := models.News{
				Title:      item.Title,
				CnTitle:    cnTitle,
				Content:    cnDescription,
				Url:        item.Link,
				SourceId:   item.GUID,
				SourceName: name,
				Image:      image,
			}

			saveNews(news)
		}
	})
}
