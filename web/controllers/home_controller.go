package controllers

import (
	"github.com/GoLangDream/iceberg/web"
	"hot_news/service"
	"hot_news/works"
)

func init() {
	web.RegisterController(HomeController{})
}

type HomeController struct {
	*web.BaseController
}

func (c *HomeController) Index() {
	str := "一款服务于 Go 开发者的依赖注入框架，方便搭建任何 Go 应用。 A Golang depenedency injection framework, helps developers to build any go application"
	c.Text(service.TranslateString(str))
}

func (c *HomeController) Update() {
	works.SyncGithubTrending()
	works.SyncHackNews()
	c.Text("success")
}
