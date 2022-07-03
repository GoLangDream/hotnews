package works

import (
	"errors"
	"github.com/GoLangDream/iceberg/database"
	"github.com/GoLangDream/iceberg/log"
	"gorm.io/gorm"
	"hot_news/models"
)

func checkNewsExists(sourceName, sourceId string) bool {
	var n models.News
	result := database.DBConn.Where(
		"source_id = ? AND source_name = ?",
		sourceId,
		sourceName,
	).First(&n)

	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return false
	}
	return true
}

func saveNews(news models.News) {
	var n models.News
	result := database.DBConn.Where(
		"source_id = ? AND source_name = ?",
		news.SourceId,
		news.SourceName,
	).First(&n)

	if !errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return
	}

	result = database.DBConn.Create(&news)

	if result.Error != nil {
		log.Infof("创建 hot_news 文章失败, 名称 [%s], %s", news.CnTitle, result.Error.Error())
		return
	}

	log.Infof("创建 hot_news 文章 [%s]", news.CnTitle)
}
