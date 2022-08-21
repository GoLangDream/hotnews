package works

import (
	"github.com/GoLangDream/iceberg/database"
	"github.com/GoLangDream/iceberg/log"
	"github.com/GoLangDream/iceberg/work"
	"hot_news/models"
	"hot_news/service/google"
	"time"
)

func init() {
	work.Register("fetch_image", "@every 10m", fetchImage)
}

func fetchImage() {
	var news []*models.News

	database.DBConn.
		Where("image = ? and created_at > ? ", "", time.Now().AddDate(0, 0, -3)).
		Limit(50).
		Order("id DESC").
		Find(&news)

	for _, m := range news {
		log.Infof("开始抓取图片 [%d], %s", m.ID, m.Title)
		m.Image = google.GetSearchImage(m.Title)
		database.DBConn.Save(m)
		time.Sleep(2 * time.Second)
	}
}
