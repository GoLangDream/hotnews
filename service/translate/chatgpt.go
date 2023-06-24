package translate

import (
	"hot_news/service/openai"
)

// ChatGPTTranslateString 翻译字符串
func ChatGPTTranslateString(q string) string {
	return openai.GetChatGPTAck("翻译下面这段你文字 \n" + q)
}
