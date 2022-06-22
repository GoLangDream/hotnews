package works

import (
	"fmt"
	"github.com/GoLangDream/iceberg/log"
	"github.com/andygrunwald/go-trending"
	"github.com/robfig/cron/v3"
	"hot_news/models"
	"hot_news/service"
)

var githubJobID cron.EntryID
var trend = trending.NewTrending()

func init() {
	githubJobID, _ = cronTask.AddFunc("@hourly", SyncGithubTrending)
}

func SyncGithubTrending() {
	insertProject(trend.GetProjects(trending.TimeToday, ""))
	insertProject(trend.GetProjects(trending.TimeToday, "go"))
	insertProject(trend.GetProjects(trending.TimeToday, "ruby"))
	insertProject(trend.GetProjects(trending.TimeToday, "kotlin"))
	insertProject(trend.GetProjects(trending.TimeToday, "java"))
	insertProject(trend.GetProjects(trending.TimeToday, "python"))

	printCronTask("GithubTrending", githubJobID)
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

		cnDescription := service.TranslateString(project.Description)
		news := models.News{
			Title:      project.Name,
			CnTitle:    fmt.Sprintf("[%s] %s", project.Language, project.Name),
			Content:    cnDescription,
			Url:        project.URL.String(),
			SourceId:   project.Name,
			SourceName: "github_trending",
		}

		saveNews(news)
	}
}
