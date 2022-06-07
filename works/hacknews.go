package works

import (
	"errors"
	"github.com/GoLangDream/iceberg/database"
	"github.com/GoLangDream/iceberg/log"
	"github.com/go-shiori/go-readability"
	"github.com/peterhellberg/hn"
	"gorm.io/gorm"
	"hot_news/models"
	"hot_news/service"
	"strconv"
	"time"
)

var hnClient = hn.DefaultClient

func init() {
	cronTask.AddFunc("@hourly", syncHotNews)
}

func syncHotNews() {

	ids, _ := hnClient.TopStories()

	log.Infof("的到 %d 篇文章", len(ids))

	for _, id := range ids {
		item, _ := hnClient.Item(id)
		var news models.News

		result := database.DBConn.Where(
			"source_id = ? AND source_name = ?",
			strconv.Itoa(item.ID),
			"hacknews",
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
			SourceName: "hacknews",
		}

		result = database.DBConn.Create(&news)

		if result.Error != nil {
			log.Infof("创建失败, 文章 [%s], %s", cnTitle, result.Error.Error())
			continue
		}

		log.Infof("创建 Hacknews 文章 [%s]", cnTitle)
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
