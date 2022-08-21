package controllers

import (
	"fmt"
	"github.com/GoLangDream/iceberg/web"
	"github.com/GoLangDream/iceberg/work"
	"hot_news/service"
	"hot_news/service/google"
	"hot_news/service/translate"
)

func init() {
	web.RegisterController(HomeController{})
}

type HomeController struct {
	*web.BaseController
}

func (c *HomeController) Index() {
	url := google.GetSearchImage("What takes years and costs $20K? A San Francisco trash can")
	c.Text(url)
}

func (c *HomeController) WebContent() {
	url := c.Query("url")
	excerpt, image, _ := service.FetchWebContent(url)
	c.Text(fmt.Sprintf("excerpt: %s\n, image: %s", excerpt, image))
}

func (c *HomeController) Translate() {
	text := c.Query("text")
	cnText, _ := translate.Content(text)
	c.Text(fmt.Sprintf("翻译的内容是 [%s]", cnText))
}

func (c *HomeController) Update() {
	name := c.Query("name")

	work.RunWorksNow(name)

	c.Text("success")
}
