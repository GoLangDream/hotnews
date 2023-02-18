package controllers

import (
	"github.com/GoLangDream/iceberg/web"
	"hot_news/models"
	"hot_news/service"
	"strconv"
)

func init() {
	web.RegisterController(NewsController{})
}

type NewsController struct {
	*web.BaseController
}

func (c *NewsController) Index() {

	lastID, _ := strconv.Atoi(c.Query("last_id"))

	news := service.LastNews(lastID)
	var newLastID uint = 0
	if len(news) != 0 {
		newLastID = news[len(news)-1].ID
	}

	c.Json(map[string]any{
		"code":    0,
		"message": "success",
		"last_id": newLastID,
		"items":   news,
	})
}

func (c *NewsController) Read() {

	newsID, _ := strconv.Atoi(c.Query("id"))

	c.DB().Model(&models.News{}).
		Where("id = ?", newsID).
		Update("is_readed", true)

	c.Json(map[string]any{
		"code":    0,
		"message": "success",
		"mews_id": newsID,
	})
}
