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
	result, err := GoogleTranslateString(q)
	if err == nil {
		return result, nil
	}

	log.Infof("google翻译错误 %d, %s", err.Code, err.Message)

	switch err.Code {
	case ErrorTypeServerError:
		return BaiduTranslateString(q), nil
	case ErrorTypeTooManyRequests:
		return BaiduTranslateString(q), nil
	case ErrorTypeResponseFormatError:
		return BaiduTranslateString(q), nil
	default:
		return "", err
	}
}
