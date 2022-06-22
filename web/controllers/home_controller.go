package controllers

import (
	"github.com/GoLangDream/iceberg/web"
	"hot_news/works"
)

func init() {
	web.RegisterController(HomeController{})
}

type HomeController struct {
	*web.BaseController
}

func (c *HomeController) Index() {
	c.Text("hello word")
}

func (c *HomeController) Update() {
	works.SyncGithubTrending()
	works.SyncHackNews()
	c.Text("success")
}
