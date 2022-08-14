package works

import (
	"github.com/GoLangDream/iceberg/database"
	"github.com/GoLangDream/iceberg/work"
	"hot_news/models"
	"time"
)

func init() {
	work.Register("translate_news", "@every 10m", translateNews)
}

func translateNews() {
	var news []*models.News

	database.DBConn.
		Where("is_translate = ?", false).
		Limit(50).
		Order("id DESC").
		Find(&news)

	for _, m := range news {
		m.Translate()
		time.Sleep(2 * time.Second)
	}
}
