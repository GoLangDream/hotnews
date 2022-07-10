package models

import "gorm.io/gorm"

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
}
