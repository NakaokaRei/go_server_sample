package model

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
	Session_id string `json:"session_id"`
	User_id int `json:"user_id"`
	Timestamp int `json:"timestamp"`
	Eating_log []EatLog `json:"eating_log"`
	Color_summary_red int `json:"color_summary_red"`
	Color_summary_yellow int `json:"color_summary_yellow"`
	Color_summary_green int `json:"color_summary_green"`
	Color_summary_purple int `json:"color_summary_purple"`
	Color_summary_brown int `json:"color_summary_brown"`
	Color_summary_white int `json:"color_summary_white"`
	Color_summary_black int `json:"color_summary_black"`
	Location_lat float64 `json:"location_lat"`
	Location_lng float64 `json:"location_lng"`
	Location_alt float64 `json:"location_alt"`
	Location_acc float64 `json:"location_acc"`
}

type EatLog struct {
	Timestamp int `json:"t"`
	Color int `json:"c"`
}

type AfterRequest struct {
	Session_id string `json:"session_id"`
	User_id int `json:"user_id"`
	Timestamp int `json:"timestamp"`
	Photo string `json:"photo:`
	Location_lat float64 `json:"location_lat"`
	Location_lng float64 `json:"location_lng"`
	Location_alt float64 `json:"location_alt"`
	Location_acc float64 `json:"location_acc"`
}

// beforeテーブル
type BeforeTable struct {
	Session_id string
    User_id int
	Timestamp int
	Photo_path string
	Location_lat float64
	Location_lng float64
	Location_alt float64
	Location_acc float64
}

// フードテーブル
type FoodTable struct {
	Session_id string
	Food_name string
}

// ログテーブル
type LogTable struct {
	Session_id string
	User_id int
	Timestamp int
	Color_summary_red int
	Color_summary_yellow int
	Color_summary_green int
	Color_summary_purple int
	Color_summary_brown int
	Color_summary_white int
	Color_summary_black int
	Location_lat float64
	Location_lng float64
	Location_alt float64
	Location_acc float64
}

// 摂食ログテーブル
type LogsTable struct {
	Session_id string
	Timestamp int
	Color int
}

// アフターテーブル
type AfterTable struct {
    Session_id string
    User_id int
	Timestamp int
	Photo_path string
	Location_lat float64
	Location_lng float64
	Location_alt float64
	Location_acc float64
}