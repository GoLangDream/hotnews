package controllers

import (
	"fmt"
	"github.com/GoLangDream/iceberg/web"
	"github.com/GoLangDream/iceberg/work"
	"hot_news/models"
	"hot_news/service"
	"strconv"
)

func init() {
	web.RegisterController(HomeController{})
}

type HomeController struct {
	*web.BaseController
}

func (c *HomeController) Index() {
	c.Text("hello world")
}

func (c *HomeController) News() {
	c.Header("Access-Control-Allow-Origin", "*")
	lastID, _ := strconv.Atoi(c.Query("last_id"))

	news := service.LastNews(lastID)

	c.Json(map[string]any{
		"code":    0,
		"message": "success",
		"last_id": news[len(news)-1].ID,
		"items":   news,
	})
}

func (c *HomeController) ReadNews() {
	c.Header("Access-Control-Allow-Origin", "*")
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

func (c *HomeController) WebContent() {
	url := c.Query("url")
	excerpt, image, _ := service.FetchWebContent(url)
	c.Text(fmt.Sprintf("excerpt: %s, image: %s", excerpt, image))
}

func (c *HomeController) Translate() {
	text := c.Query("text")
	c.Text("翻译的内容是" + service.BaiduTranslateString(text))
}

func (c *HomeController) Update() {
	name := c.Query("name")

	work.RunWorksNow(name)

	c.Text("success")
}
