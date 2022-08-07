package service

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"github.com/GoLangDream/iceberg/log"
	"github.com/gookit/config/v2"
	"github.com/imroc/req/v3"
	"math/rand"
	"strconv"
	"time"
)

type baiduResponse struct {
	From        string             `json:"from"`
	To          string             `json:"to"`
	TransResult []baiduTransResult `json:"trans_result"`
	ErrorCode   string             `json:"error_code"`
}

type baiduTransResult struct {
	Src string `json:"src"`
	Dst string `json:"dst"`
}

func baiduRequest(q string) map[string]string {
	rand.Seed(time.Now().UnixNano())

	appid := config.String("application.baidu.appID")
	salt := strconv.Itoa(rand.Int())
	signContent := fmt.Sprintf("%s%s%s%s", appid, q, salt, config.String("application.baidu.appSecret"))
	sign := md5.Sum([]byte(signContent))
	return map[string]string{
		"q":     q,
		"from":  "auto",
		"to":    "zh",
		"appid": appid,
		"salt":  salt,
		"sign":  hex.EncodeToString(sign[:]),
	}
}

// BaiduTranslateString 翻译字符串
func BaiduTranslateString(q string) string {
	url := "https://fanyi-api.baidu.com/api/trans/vip/translate"
	client := req.C()
	var response baiduResponse

	rep, err := client.R().
		SetFormData(baiduRequest(q)).
		SetResult(&response).
		Post(url)

	if err == nil && rep.IsSuccess() {
		if response.ErrorCode != "" {
			log.Infof("翻译字符串错误 [%s]", response.ErrorCode)
		}
		return response.TransResult[0].Dst
	}

	return ""
}
