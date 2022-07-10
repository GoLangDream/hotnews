package service

import (
	"github.com/GoLangDream/iceberg/database"
	"hot_news/models"
)

type News struct {
	ID      uint   `json:"id"`
	Title   string `json:"title"`
	Content string `json:"content"`
	Image   string `json:"image"`
}

func LastNews(id int) []News {
	var _news []models.News
	var news []News

	database.DBConn.
		Debug().
		Limit(10).
		Order("id DESC").
		Where("id > ?", id).
		Find(&_news)

	for _, m := range _news {
		news = append(news, News{
			ID:      m.ID,
			Title:   m.CnTitle,
			Content: m.Content,
			Image:   m.Image,
		})
	}

	return news
}
