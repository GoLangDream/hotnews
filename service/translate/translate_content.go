package translate

import "github.com/GoLangDream/iceberg/log"

type Error struct {
	Code    ErrorType
	Message string
}

type ErrorType int

const (
	ErrorTypeUnknown ErrorType = iota
	ErrorTypeTooManyRequests
	ErrorTypeServerError
	ErrorTypeResponseFormatError
	ErrorTypeOther
)

func Content(q string) (string, *Error) {
	result, error := GoogleTranslateString(q)
	if error == nil {
		return result, nil
	}

	log.Infof("google翻译错误 %d, %s", error.Code, error.Message)

	switch error.Code {
	case ErrorTypeServerError:
		return BaiduTranslateString(q), nil
	case ErrorTypeTooManyRequests:
		return BaiduTranslateString(q), nil
	case ErrorTypeResponseFormatError:
		return BaiduTranslateString(q), nil
	default:
		return "", error
	}
}
