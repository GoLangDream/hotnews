package service

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
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

func TranslateString(q string) string {
	url := "https://openapi.youdao.com/api"
	client := req.C()
	var result ydResult

	rep, err := client.R().
		SetFormData(createRepo(q)).
		SetResult(&result).
		Post(url)

	if err != nil && rep.IsSuccess() {
		return result.Translation[0]
	}

	return ""
}
