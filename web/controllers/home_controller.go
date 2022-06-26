package controllers

import (
	"github.com/GoLangDream/iceberg/web"
	"github.com/GoLangDream/iceberg/work"
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
	work.RunWorkNow("github_trending")
	work.RunWorkNow("hacknews")
	c.Text("success")
}
