package translate

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"github.com/GoLangDream/iceberg/log"
	"github.com/forPelevin/gomoji"
	"github.com/google/uuid"
	"github.com/gookit/config/v2"
	"github.com/imroc/req/v3"
	"time"
)

type ydResult struct {
	ErrorCode   string   `json:"errorCode"`
	Query       string   `json:"query"`
	Translation []string `json:"translation"`
	Basic       string   `json:"basic"`
	//Web          []map[string][]string `json:"web"`
	L            string   `json:"l"`
	Dict         string   `json:"dict"`
	Webdict      string   `json:"webdict"`
	TSpeakUrl    string   `json:"tSpeakUrl"`
	SpeakUrl     string   `json:"speakUrl"`
	ReturnPhrase []string `json:"returnPhrase"`
}

type ydHtmlResult struct {
	Data         string `json:"data"`
	ErrorCode    string `json:"errorCode"`
	ErrorMessage string `json:"errorMessage"`
}

func createRepo(q string) map[string]string {
	appKey := config.String("application.youdao.appKey")
	salt := uuid.New().String()
	curtime := fmt.Sprintf("%d", time.Now().Unix())
	signContent := fmt.Sprintf("%s%s%s%s%s", appKey, input(q), salt, curtime, config.String("application.youdao.appSecret"))
	sign := sha256.Sum256([]byte(signContent))
	return map[string]string{
		"q":        q,
		"from":     "en",
		"to":       "zh-CHS",
		"appKey":   appKey,
		"salt":     salt,
		"sign":     hex.EncodeToString(sign[:]),
		"signType": "v3",
		"curtime":  curtime,
	}
}

func input(q string) string {
	str := []rune(q)
	length := len(str)
	if length <= 20 {
		return q
	}
	return fmt.Sprintf("%s%d%s", string(str[0:10]), length, string(str[length-10:length]))
}

// YDTranslateString 翻译的内容，暂时不支持 emoji, 已经提交给 youdao，等待他们处理
func YDTranslateString(q string) string {
	url := "https://openapi.youdao.com/api"
	client := req.C()
	var result ydResult

	rep, err := client.R().
		SetFormData(createRepo(gomoji.RemoveEmojis(q))).
		SetResult(&result).
		Post(url)

	if err == nil && rep.IsSuccess() {
		if result.ErrorCode != "0" {
			log.Infof("翻译字符串错误 [%s]", result.ErrorCode)
		}
		return result.Translation[0]
	}

	return ""
}

// YDTranslateHtml 翻译的内容，暂时不支持 emoji, 已经提交给 youdao，等待他们处理
func YDTranslateHtml(q string) string {
	url := "https://openapi.youdao.com/translate_html"
	client := req.C().DevMode()
	var result ydHtmlResult

	rep, err := client.R().
		SetFormData(createRepo(gomoji.RemoveEmojis(q))).
		SetResult(&result).
		Post(url)

	if err == nil && rep.IsSuccess() {
		if result.ErrorCode != "0" {
			log.Infof("翻译html错误 [%s] %s", result.ErrorCode, result.ErrorMessage)
		}
		return result.Data
	}

	return ""
}
