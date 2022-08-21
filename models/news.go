package models

import (
	"errors"
	"github.com/GoLangDream/iceberg/database"
	"github.com/GoLangDream/iceberg/log"
	"gorm.io/gorm"
	"hot_news/service/rss"
	"hot_news/service/translate"
	"strings"
)

type News struct {
	gorm.Model
	Title       string
	CnTitle     string
	Content     string
	Url         string
	SourceId    string
	SourceName  string
	IsSlackSend bool
	Image       string
	IsTranslate bool
	CnContent   string
}

func (news *News) BeforeCreate(tx *gorm.DB) (err error) {
	if news.Exists() {
		return errors.New("记录已经存在, 不能保存")
	}
	return nil
}

func (news *News) Exists() bool {
	var n News
	result := database.DBConn.Where(
		"source_id = ? AND source_name = ?",
		news.SourceId,
		news.SourceName,
	).First(&n)

	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return false
	}
	return true
}

func (news *News) Translate() {
	if news.IsTranslate || !news.NeedTranslate() {
		return
	}
	log.Infof("开始翻译文章 [%d], %s", news.ID, news.Title)

	cnTitle, errTitle := translate.Content(news.Title)
	if errTitle != nil {
		return
	}

	//cnContent, errContent := translate.Content(news.LessContent())
	//
	//if errContent == nil {
	//	news.CnContent = cnContent
	//} else {
	//	log.Infof("google 翻译错误 [%d] %s", errContent.Code, errContent.Message)
	//}

	news.CnTitle = cnTitle

	news.IsTranslate = true
	database.DBConn.Save(news)
}

func (news *News) NeedTranslate() bool {
	if news.SourceName == "hacknews" || news.SourceName == "github_trending" {
		return true
	}
	return rss.Sources[news.SourceName].NeedTranslate
}

func (news *News) ShowTitle() string {
	if news.NeedTranslate() && news.IsTranslate && news.CnTitle != "" {
		return news.CnTitle
	}
	return news.Title
}

func (news *News) ShowContent() string {
	if news.NeedTranslate() && news.IsTranslate && news.CnContent != "" {
		return news.CnContent
	}
	return news.LessContent()
}

func (news *News) LessContent() string {
	//content := ""
	//
	//for _, line := range strings.Split(strings.TrimSpace(news.Content), "\n") {
	//	if strings.TrimSpace(line) != "" {
	//		content = line
	//		break
	//	}
	//}

	// 由于内容字段太长，默认只用显示前 1000 个字符
	tmp := []rune(strings.TrimSpace(news.Content))
	contentLength := len(tmp)
	if contentLength > 1000 {
		contentLength = 1000
	}
	return string(tmp[:contentLength])
}

func (news *News) fetchGoogleSearchImage() {
	if strings.TrimSpace(news.Image) == "" {

	}
}
