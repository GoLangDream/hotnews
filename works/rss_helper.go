package works

import (
	"github.com/GoLangDream/iceberg/log"
	"github.com/GoLangDream/iceberg/work"
	"github.com/mmcdole/gofeed"
	"hot_news/models"
	"hot_news/service"
)

var fp = gofeed.NewParser()

func syncRss(name, url string, needTranslate ...bool) {

	_needTranslate := true

	if len(needTranslate) > 0 && needTranslate[0] == false {
		_needTranslate = false
	}

	work.Register(name, "@hourly", func() {
		feed, err := fp.ParseURL(url)
		if err != nil {
			log.Infof("获取 %s 的rss地址错误 %s", name, err)
			return
		}

		for _, item := range feed.Items {
			if checkNewsExists(name, item.GUID) {
				continue
			}

			cnTitle := item.Title

			if _needTranslate {
				cnTitle = service.TranslateString(item.Title)
			}

			if len(item.Link) > 250 {
				log.Infof("文章 [%s] 的 url [%s] 超长", cnTitle, item.Link)
				continue
			}

			news := models.News{
				Title:      item.Title,
				CnTitle:    cnTitle,
				Content:    "",
				Url:        item.Link,
				SourceId:   item.GUID,
				SourceName: name,
			}

			saveNews(news)
		}
	})
}
