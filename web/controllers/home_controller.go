package controllers

import (
	"github.com/GoLangDream/iceberg/web"
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
	c.Text("success")
}
