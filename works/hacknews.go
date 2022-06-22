package works

import (
	"github.com/GoLangDream/iceberg/log"
	"github.com/peterhellberg/hn"
	"github.com/robfig/cron/v3"
	"hot_news/models"
	"hot_news/service"
	"strconv"
)

var hnClient = hn.DefaultClient
var hacknewsJobID cron.EntryID
var hacknewsSourceName = "hacknews"

func init() {
	hacknewsJobID, _ = cronTask.AddFunc("@hourly", SyncHackNews)
}

func SyncHackNews() {

	ids, _ := hnClient.TopStories()

	log.Infof("的到 %d 篇文章", len(ids))

	for _, id := range ids {
		item, err := hnClient.Item(id)

		if err != nil {
			log.Infof("获取 %s 的文章失败", err)
			continue
		}

		if checkNewsExists(hacknewsSourceName, strconv.Itoa(item.ID)) {
			continue
		}

		cnTitle := service.TranslateString(item.Title)

		news := models.News{
			Title:      item.Title,
			CnTitle:    cnTitle,
			Content:    "",
			Url:        item.URL,
			SourceId:   strconv.Itoa(item.ID),
			SourceName: hacknewsSourceName,
		}

		saveNews(news)

		printCronTask("hacknews", hacknewsJobID)
	}

}
