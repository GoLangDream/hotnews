package works

import (
	"fmt"
	"github.com/GoLangDream/iceberg/log"
	"github.com/GoLangDream/iceberg/work"
	"github.com/andygrunwald/go-trending"
	"hot_news/models"
	"hot_news/service"
)

var trend = trending.NewTrending()

func init() {
	work.Register("github_trending", "@hourly", syncGithubTrending)
}

func syncGithubTrending() {
	insertProject(trend.GetProjects(trending.TimeToday, ""))
	insertProject(trend.GetProjects(trending.TimeToday, "go"))
	insertProject(trend.GetProjects(trending.TimeToday, "ruby"))
	insertProject(trend.GetProjects(trending.TimeToday, "kotlin"))
	insertProject(trend.GetProjects(trending.TimeToday, "java"))
	insertProject(trend.GetProjects(trending.TimeToday, "python"))

}

func insertProject(projects []trending.Project, err error) {
	if err != nil {
		log.Infof("获取github trend失败, %s", err)
		return
	}

	for _, project := range projects {
		_, image, _ := service.FetchWebContent(project.URL.String())

		news := &models.News{
			Title:       fmt.Sprintf("[%s] %s", project.Language, project.Name),
			Content:     project.Description,
			Url:         project.URL.String(),
			SourceId:    project.Name,
			SourceName:  "github_trending",
			Image:       image,
			IsTranslate: false,
		}

		saveNews(news)
	}
}
