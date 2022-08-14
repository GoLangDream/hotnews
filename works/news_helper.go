package works

import (
	"github.com/GoLangDream/iceberg/database"
	"github.com/GoLangDream/iceberg/log"
	"hot_news/models"
)

func saveNews(news *models.News) {
	result := database.DBConn.Create(&news)

	if result.Error != nil {
		if result.Error.Error() != "记录已经存在, 不能保存" {
			log.Infof("创建 hot_news 文章失败, 名称 [%s], %s", news.ShowTitle(), result.Error.Error())
		}
		return
	}

	log.Infof("创建 hot_news 文章 [%s]", news.ShowTitle())
}
