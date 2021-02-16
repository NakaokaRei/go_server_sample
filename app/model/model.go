package model

import "github.com/jinzhu/gorm"

type MessageModel struct {
	Message string `json:"message"`
	Title string `json:"title"`
}

type BeforeRequest struct {
	User_id int `json:"user_id"`
	Timestamp int `json:"timestamp"`
	Photo string `json:"photo:`
	Food_name []string `json:"food_name"`
	Location_lat float64 `json:"location_lat"`
	Location_lng float64 `json:"location_lng"`
	Location_alt float64 `json:"location_alt"`
	Location_acc float64 `json:"location_acc"`
}

type LogRequest struct {
	User_id int `json:"user_id"`
	Timestamp int `json:"timestamp"`
	Eating_log string `json:"eating_log"`
	Color_summary []int `json:"color_summary"`
	Location_lat float64 `json:`
	Location_lng float64 `json:`
	Location_alt float64 `json:`
	Location_acc float64 `json:`
}

type AfterRequest struct {
	User_id int `json:"user_id"`
	Timestamp int `json:"timestamp"`
	Photo string `json:"photo:`
	Location_lat float64 `json:`
	Location_lng float64 `json:`
	Location_alt float64 `json:`
	Location_acc float64 `json:`
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