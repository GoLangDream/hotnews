package controllers

import (
	"github.com/GoLangDream/iceberg/web"
	"github.com/gorilla/feeds"
	"time"
)

func init() {
	web.RegisterController(HomeController{})
}

type HomeController struct {
	*web.BaseController
}

func (c *HomeController) Index() {
	c.Text("hello")
}

func getFeed() string {
	now := time.Now()
	feed := &feeds.Feed{
		Title:       "Hot News",
		Link:        &feeds.Link{Href: "http://localhost:3000/"},
		Description: "一些热门的新闻",
		Author:      &feeds.Author{Name: "jimxl", Email: "tianxiaxl@gmail.com"},
		Created:     now,
	}

	feed.Items = []*feeds.Item{
		{
			Title:       "Limiting Concurrency in Go",
			Link:        &feeds.Link{Href: "http://jmoiron.net/blog/limiting-concurrency-in-go/"},
			Description: "A discussion on controlled parallelism in golang",
			Author:      &feeds.Author{Name: "Jason Moiron", Email: "jmoiron@jmoiron.net"},
			Created:     now,
		},
		{
			Title:       "Logic-less Template Redux",
			Link:        &feeds.Link{Href: "http://jmoiron.net/blog/logicless-template-redux/"},
			Description: "More thoughts on logicless templates",
			Created:     now,
		},
		{
			Title:       "Idiomatic Code Reuse in Go",
			Link:        &feeds.Link{Href: "http://jmoiron.net/blog/idiomatic-code-reuse-in-go/"},
			Description: "How to use interfaces <em>effectively</em>",
			Created:     now,
		},
	}
	rss, _ := feed.ToRss()
	return rss
}
