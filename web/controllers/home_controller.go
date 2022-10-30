package controllers

import (
	"fmt"
	"github.com/GoLangDream/iceberg/environment"
	"github.com/GoLangDream/iceberg/log"
	"github.com/GoLangDream/iceberg/web"
	"github.com/GoLangDream/iceberg/work"
	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"hot_news/service"
	"hot_news/service/aliyun"
	"hot_news/service/google"
	"hot_news/service/translate"
	"math/rand"
	"net/http"
	"path"
	"time"
)

func init() {
	web.RegisterController(HomeController{})
}

type HomeController struct {
	*web.BaseController
}

func (c *HomeController) Index() {
	url := google.GetSearchImage("What takes years and costs $20K? A San Francisco trash can")
	c.Text(url)
}

func (c *HomeController) WebContent() {
	url := c.Query("url")
	excerpt, image, _ := service.FetchWebContent(url)
	c.Text(fmt.Sprintf("excerpt: %s\n, image: %s", excerpt, image))
}

func (c *HomeController) Translate() {
	text := c.Query("text")
	cnText, _ := translate.Content(text)
	c.Text(fmt.Sprintf("翻译的内容是 [%s]", cnText))
}

func (c *HomeController) Update() {
	name := c.Query("name")

	work.RunWorksNow(name)

	c.Text("success")
}

func (c *HomeController) GetImageUploadUrl() {

	response := aliyun.GetAssumeRole()

	ossClient, err := oss.New(
		"https://oss-cn-shenzhen.aliyuncs.com",
		response.Credentials.AccessKeyId,
		response.Credentials.AccessKeySecret,
		oss.SecurityToken(response.Credentials.SecurityToken),
	)

	bucket, err := ossClient.Bucket("dreamspace-app")

	fileName := c.Query("file_name")
	filesuffix := path.Ext(fileName)

	ossFileName := ""

	if environment.IsProduction() {
		ossFileName = fmt.Sprintf("images/%s-%d%s", time.Now().Format("20060102-150405"), rand.Intn(100), filesuffix)
	} else {
		ossFileName = fmt.Sprintf("images_test/%s-%d%s", time.Now().Format("20060102-150405"), rand.Intn(100), filesuffix)
	}

	options := []oss.Option{
		oss.ContentType(c.Query("content_type")),
	}

	signedURL, err := bucket.SignURL(ossFileName, oss.HTTPPut, 300, options...)
	if err != nil {
		log.Infof("获取上传 URL 错误 %s", err)
	}
	c.Json(map[string]any{
		"upload_url": signedURL,
		"file_url":   ossFileName,
	})
}

func (c *HomeController) ImageUrl() {

	response := aliyun.GetAssumeRole()

	ossClient, err := oss.New(
		"https://oss-cn-shenzhen.aliyuncs.com",
		response.Credentials.AccessKeyId,
		response.Credentials.AccessKeySecret,
		oss.SecurityToken(response.Credentials.SecurityToken),
	)

	bucket, err := ossClient.Bucket("dreamspace-app")

	fileName := c.Query("file_name")

	signedURL, err := bucket.SignURL(fileName, oss.HTTPGet, 60*3)

	if err != nil {
		log.Infof("获取下载 URL 错误 %s", err)
	}

	c.RedirectTo(signedURL, http.StatusFound)

}
