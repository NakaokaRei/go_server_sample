package model

import "github.com/jinzhu/gorm"

type MessageModel struct {
	Message string `json:"message"`
	Title string `json:"title"`
}

type BeforeModel struct {
    gorm.Model
    Text   string
    Status string
}

type LogModel struct {
    gorm.Model
    Text   string
    Log string
}

type AfterModel struct {
    gorm.Model
    Text   string
    AfterPath string
}