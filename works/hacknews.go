package works

import (
	"github.com/GoLangDream/iceberg/log"
	"github.com/GoLangDream/iceberg/work"
	"github.com/peterhellberg/hn"
	"hot_news/models"
	"hot_news/service"
	"strconv"
)

var hnClient = hn.DefaultClient
var hacknewsSourceName = "hacknews"

func init() {
	work.Register(hacknewsSourceName, "@hourly", syncHackNews)
}

func syncHackNews() {

	ids, _ := hnClient.TopStories()

	log.Infof("的到 %d 篇文章", len(ids))

	for _, id := range ids {
		item, err := hnClient.Item(id)

		if err != nil {
			log.Infof("获取 %s 的文章失败", err)
			continue
		}

		excerpt, image, _ := service.FetchWebContent(item.URL)

		news := &models.News{
			Title:       item.Title,
			Content:     excerpt,
			Url:         item.URL,
			SourceId:    strconv.Itoa(item.ID),
			SourceName:  hacknewsSourceName,
			Image:       image,
			IsTranslate: false,
		}

		saveNews(news)
	}
}
