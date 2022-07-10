package controllers

import (
	"fmt"
	"github.com/GoLangDream/iceberg/web"
	"github.com/GoLangDream/iceberg/work"
	"hot_news/service"
)

func init() {
	web.RegisterController(HomeController{})
}

type HomeController struct {
	*web.BaseController
}

func (c *HomeController) Index() {
	excerpt, image, _ := service.FetchWebContent("https://www.producthunt.com/posts/amplify-ui")
	cnExcerpt := service.TranslateString(excerpt)
	c.Text(fmt.Sprintf("excerpt: %s, image: %s, cn excerpt: %s", excerpt, image, cnExcerpt))
}

func (c *HomeController) Update() {
	name := c.Query("name")

	work.RunWorksNow(name)

	c.Text("success")
}
