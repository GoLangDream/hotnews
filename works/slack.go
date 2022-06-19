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
	slackJobID, _ = cronTask.AddFunc("@every 1m", SendHotNewsToSlack)
}

func SendHotNewsToSlack() {
	log.Infof("开始检测新文章")
	initSlackConfig()

	var news []models.News
	database.DBConn.Debug().Limit(100).
		Order("id").
		Where("is_slack_send = ?", false).
		Find(&news)

	sendSlackMessage(news)

	printCronTask("slack", slackJobID)
}

func initSlackConfig() {
	slackApi = slack.New(
		config.String("application.slack.token"),
		slack.OptionDebug(environment.IsDevelopment()),
	)

	hotNewsChannelID = config.String("application.slack.hotNewsChannelID")
}

func sendSlackMessage(news []models.News) {
	newsCount := len(news)
	if newsCount <= 0 {
		return
	}

	for _, m := range news {
		log.Infof("发送文章 %s", m.CnTitle)
		newsSection := newTextSection(fmt.Sprintf("<%s|%s>", m.Url, m.CnTitle))
		infoSection := newsInfoSection(m)

		_, _, err := slackApi.PostMessage(
			hotNewsChannelID,
			slack.MsgOptionBlocks(slack.NewDividerBlock(), infoSection, newsSection, slack.NewDividerBlock()),
		)

		if err != nil {
			log.Infof("发送消息错误, 原因: [%s]", err)
			return
		}

		database.DBConn.Model(&m).Update("is_slack_send", true)
		time.Sleep(time.Duration(1) * time.Second)
	}

}

func newTextSection(content string) *slack.SectionBlock {
	block := slack.NewTextBlockObject(
		"mrkdwn",
		content,
		false,
		false,
	)
	return slack.NewSectionBlock(block, nil, nil)
}

func newsInfoSection(news models.News) *slack.SectionBlock {
	return slack.NewSectionBlock(
		nil,
		[]*slack.TextBlockObject{
			{
				Type: "mrkdwn",
				Text: fmt.Sprintf("*【文章编号】:* %d\n*【文章来源】:* %s\n*【爬取时间】:* %s\n", news.ID, news.SourceName, news.CreatedAt.Format("01-02 15:04")),
			},
		},
		nil,
	)
}
