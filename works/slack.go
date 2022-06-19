package works

import (
	"fmt"
	"github.com/GoLangDream/iceberg/database"
	"github.com/GoLangDream/iceberg/environment"
	"github.com/GoLangDream/iceberg/log"
	"github.com/gookit/config/v2"
	"github.com/robfig/cron/v3"
	"github.com/slack-go/slack"
	"hot_news/models"
	"time"
)

var (
	slackApi         *slack.Client
	hotNewsChannelID string
	slackJobID       cron.EntryID
)

func init() {
	slackJobID, _ = cronTask.AddFunc("@every 10m", SendHotNewsToSlack)
}

func SendHotNewsToSlack() {
	log.Infof("开始检测新文章")
	initSlackConfig()

	var news []models.News
	database.DBConn.Debug().Limit(1000).
		Order("id").
		//Where("created_at > ?", time.Now().Add(-10*time.Minute)).
		Where("is_slack_send = ?", false).
		Find(&news)

	for _, m := range news {
		sendSlackMessage(
			m.SourceId,
			fmt.Sprintf("[%s] %s", m.SourceName, m.CnTitle),
			m.Url,
		)
		database.DBConn.Debug().Model(&m).Update("is_slack_send", true)
		time.Sleep(time.Second)
	}

	printCronTask("slack", slackJobID)
}

func initSlackConfig() {
	slackApi = slack.New(
		config.String("application.slack.token"),
		slack.OptionDebug(environment.IsDevelopment()),
	)

	hotNewsChannelID = config.String("application.slack.hotNewsChannelID")
}

func sendSlackMessage(id, title, url string) {
	log.Infof("发送文章 %s", title)

	approvalText := slack.NewTextBlockObject(
		"mrkdwn",
		title,
		false,
		false,
	)

	newsTextBlock := slack.NewTextBlockObject(
		"plain_text",
		"阅读详情",
		true,
		false,
	)

	newsBtn := slack.NewButtonBlockElement(
		"hotnews_id",
		id,
		newsTextBlock,
	)

	newsBtn.URL = url

	fieldsSection := slack.NewSectionBlock(
		approvalText,
		nil,
		slack.NewAccessory(newsBtn),
	)

	_, _, err := slackApi.PostMessage(
		hotNewsChannelID,
		slack.MsgOptionBlocks(fieldsSection),
		slack.MsgOptionAsUser(true),
	)
	if err != nil {
		log.Infof("发送消息错误, 原因: [%s]", err)
		return
	}
}
