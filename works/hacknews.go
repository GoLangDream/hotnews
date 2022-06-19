package works

import (
	"errors"
	"github.com/GoLangDream/iceberg/database"
	"github.com/GoLangDream/iceberg/log"
	"github.com/go-shiori/go-readability"
	"github.com/peterhellberg/hn"
	"github.com/robfig/cron/v3"
	"gorm.io/gorm"
	"hot_news/models"
	"hot_news/service"
	"strconv"
	"time"
)

var hnClient = hn.DefaultClient
var hacknewsJobID cron.EntryID
var hacknewsSourceName = "hacknews"

func init() {
	hacknewsJobID, _ = cronTask.AddFunc("@hourly", syncHackNews)
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

		var news models.News

		result := database.DBConn.Where(
			"source_id = ? AND source_name = ?",
			strconv.Itoa(item.ID),
			hacknewsSourceName,
		).First(&news)

		if !errors.Is(result.Error, gorm.ErrRecordNotFound) {
			continue
		}

		cnTitle := service.TranslateString(item.Title)

		news = models.News{
			Title:      item.Title,
			CnTitle:    cnTitle,
			Content:    getContent(item.URL),
			Url:        item.URL,
			SourceId:   strconv.Itoa(item.ID),
			SourceName: hacknewsSourceName,
		}

		result = database.DBConn.Create(&news)

		if result.Error != nil {
			log.Infof("创建失败, 文章 [%s], %s", cnTitle, result.Error.Error())
			continue
		}

		log.Infof("创建 Hacknews 文章 [%s]", cnTitle)

		printCronTask("hacknews", hacknewsJobID)
	}

}

func getContent(url string) string {
	article, err := readability.FromURL(url, 30*time.Second)
	if err != nil {
		log.Info(err)
		return ""
	}
	return article.Content
}
