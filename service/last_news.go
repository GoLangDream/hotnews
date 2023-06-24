package service

import (
	"github.com/GoLangDream/iceberg/database"
	"hot_news/models"
	"strings"
	"time"
)

type NewsResponse struct {
	ID            uint   `json:"id"`
	Title         string `json:"title"`
	Content       string `json:"content"`
	Image         string `json:"image"`
	Url           string `json:"url"`
	Source        string `json:"source"`
	NeedTranslate bool   `json:"need_translate"`
	CreatedAt     string `json:"created_at"`
}

func LastNews(id int) []*NewsResponse {
	var news []*models.News
	var newsResponses []*NewsResponse

	db := database.DBConn.
		Where("is_readed = ? and is_translate = ? and image != ''", false, true).
		Limit(10).
		Order("id DESC")

	if id <= 0 {
		db.Find(&news)
	} else {
		db.Where("id < ?", id).Find(&news)
	}

	loc, _ := time.LoadLocation("Asia/Shanghai")

	// 暂时不返回文章的内容，因为目前的排版还不好
	for _, m := range news {
		m.Translate()
		newsResponses = append(newsResponses, &NewsResponse{
			ID:            m.ID,
			Title:         m.ShowTitle(),
			Content:       m.ShowContent(),
			Image:         url(m.Image),
			Url:           m.Url,
			Source:        m.SourceName,
			NeedTranslate: m.NeedTranslate(),
			CreatedAt:     m.CreatedAt.In(loc).Format("01-02 15:04"),
		})
	}

	return newsResponses
}

func url(path string) string {
	if strings.HasPrefix(path, "//") {
		return "https:" + path
	}
	return path
}
