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
		if checkNewsExists("github_trending", project.Name) {
			continue
		}

		cnDescription := ""
		excerpt, image, err := service.FetchWebContent(project.URL.String())
		if err != nil {
			cnDescription = service.TranslateString(excerpt)
		}

		news := models.News{
			Title:      project.Name,
			CnTitle:    fmt.Sprintf("[%s] %s", project.Language, project.Name),
			Content:    cnDescription,
			Url:        project.URL.String(),
			SourceId:   project.Name,
			SourceName: "github_trending",
			Image:      image,
		}

		saveNews(news)
	}
}
