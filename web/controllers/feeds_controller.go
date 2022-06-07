package controllers

import (
	"github.com/GoLangDream/iceberg/web"
	"github.com/gorilla/feeds"
	"hot_news/models"
	"time"
)

func init() {
	web.RegisterController(FeedsController{})
}

type FeedsController struct {
	*web.BaseController
}

func (c *FeedsController) Index() {
	c.Text(c.getFeed())
}

func (c *FeedsController) getFeed() string {
	now := time.Now()
	feed := &feeds.Feed{
		Title:       "Hot News",
		Link:        &feeds.Link{Href: "http://localhost:3000/"},
		Description: "一些热门的新闻",
		Author:      &feeds.Author{Name: "jimxl", Email: "tianxiaxl@gmail.com"},
		Created:     now,
	}
	var news []models.News
	var feedItems []*feeds.Item
	c.DB().Limit(1000).
		Find(&news).
		Order("id").
		Where("created_at > ?", time.Now().Add(-24*time.Hour))

	for _, m := range news {
		feedItems = append(feedItems, &feeds.Item{
			Title:       m.CnTitle,
			Link:        &feeds.Link{Href: m.Url},
			Description: m.Content,
			Created:     m.CreatedAt,
		})
	}

	feed.Items = feedItems
	rss, _ := feed.ToRss()
	return rss
}
