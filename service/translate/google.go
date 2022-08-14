package translate

import (
	"encoding/json"
	"fmt"
	"github.com/GoLangDream/iceberg/log"
	"github.com/imroc/req/v3"
)

var googleTranslateLength int64 = 0

func GoogleTranslateString(q string) (string, *Error) {
	googleTranslateLength += int64(len(q))

	log.Infof("调用 google 翻译长度 [%d]", googleTranslateLength)

	resultStr := ""

	url := "https://translate.google.com/translate_a/single"
	client := req.C()

	rep, err := client.R().
		SetQueryParam("client", "at").
		SetQueryParam("sl", "auto").
		SetQueryParam("tl", "zh-CN").
		SetQueryParam("dt", "t").
		SetQueryParam("q", q).
		Get(url)

	if err == nil && rep.IsSuccess() {
		var tmp []json.RawMessage
		err := json.Unmarshal(rep.Bytes(), &tmp)
		if err != nil || len(tmp) == 0 {
			return "", &Error{
				ErrorTypeResponseFormatError,
				"返回内容格式错误",
			}
		}

		err = json.Unmarshal(tmp[0], &tmp)
		if err != nil || len(tmp) == 0 {
			return "", &Error{
				ErrorTypeResponseFormatError,
				"返回内容格式错误",
			}
		}

		for _, content := range tmp {
			var result []any
			err = json.Unmarshal(content, &result)
			if err != nil || len(result) == 0 {
				continue
			}
			resultStr += result[0].(string)
		}

		return resultStr, nil
	} else {
		if rep.Status == "429 Too Many Requests" {
			return "", &Error{
				ErrorTypeTooManyRequests,
				"翻译超过频次限制",
			}
		}
		return "", &Error{
			ErrorTypeServerError,
			fmt.Sprintf("翻译错误 %s", rep.Status),
		}
	}
}
