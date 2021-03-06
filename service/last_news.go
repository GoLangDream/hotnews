package service

import (
	"github.com/GoLangDream/iceberg/database"
	"hot_news/models"
	"strings"
	"time"
)

type News struct {
	ID            uint   `json:"id"`
	Title         string `json:"title"`
	Content       string `json:"content"`
	Image         string `json:"image"`
	Url           string `json:"url"`
	Source        string `json:"source"`
	NeedTranslate bool   `json:"need_translate"`
	CreatedAt     string `json:"created_at"`
}

func LastNews(id int) []News {
	var _news []models.News
	var news []News

	db := database.DBConn.
		Where("is_readed = ?", false).
		Debug().
		Limit(10).
		Order("id DESC")
	if id <= 0 {
		db.Find(&_news)
	} else {
		db.Where("id < ?", id).Find(&_news)
	}

	loc, _ := time.LoadLocation("Asia/Shanghai")

	for _, m := range _news {
		news = append(news, News{
			ID:            m.ID,
			Title:         m.CnTitle,
			Content:       m.Content,
			Image:         url(m.Image),
			Url:           m.Url,
			Source:        m.SourceName,
			NeedTranslate: m.CnTitle != m.Title,
			CreatedAt:     m.CreatedAt.In(loc).Format("01-02 15:04"),
		})
	}

	return news
}

func url(path string) string {
	if strings.HasPrefix(path, "//") {
		return "https:" + path
	}
	return path
}
