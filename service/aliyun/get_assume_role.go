package aliyun

import (
	"github.com/GoLangDream/iceberg/log"
	"github.com/aliyun/alibaba-cloud-sdk-go/services/sts"
	"github.com/gookit/config/v2"
)

func GetAssumeRole() *sts.AssumeRoleResponse {
	client, err := sts.NewClientWithAccessKey(
		"cn-shenzhen",
		config.String("application.aliyun.accessKeyID"),
		config.String("application.aliyun.accessKeySecret"),
	)

	//构建请求对象。
	request := sts.CreateAssumeRoleRequest()
	request.Scheme = "https"

	//设置参数。
	request.RoleArn = config.String("application.aliyun.arn")
	request.RoleSessionName = "upload-image"

	//发起请求，并得到响应。
	response, err := client.AssumeRole(request)

	if err != nil {
		log.Infof("获取 sts 错误 %s", err)
	}

	return response
}
