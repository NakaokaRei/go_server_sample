package main

import(
    "github.com/gin-gonic/gin"
    "github.com/jinzhu/gorm"
    _ "github.com/google/uuid"
    "./model"
    _ "github.com/mattn/go-sqlite3"
    //"net/http"
)

func main() {
	r := gin.Default()

    dbInit()
    
	r.GET("/helloworld", helloWorld)
    r.POST("/message", uploadMessage)
    r.POST("/before", breforeResponse)
    r.POST("/log", logResponse)
    r.POST("/after", afterResponse)
	r.Run()
}

//DB初期化
func dbInit() {
    db, err := gorm.Open("sqlite3", "model/database.sqlite3")
    if err != nil {
        panic("データベース開けず！（dbInit）")
    }
    db.AutoMigrate(&model.BeforeTable{})
    db.AutoMigrate(&model.FoodTable{})
    db.AutoMigrate(&model.LogTable{})
    db.AutoMigrate(&model.LogsTable{})
    db.AutoMigrate(&model.AfterTable{})
    defer db.Close()
}

//CREATE
func dbInsert(text string, status string) {
    db, err := gorm.Open("sqlite3", "model/database.sqlite3")
    if err != nil {
        panic("データベース開けず！（dbInsert)")
    }
    db.Create(&model.FoodTable{Session_id: text, Food_name: status})
    defer db.Close()
}

//getメゾットのチュートリアル
func helloWorld(c *gin.Context) {
    c.JSON(200, gin.H{
        "message": "helloworld",
    })
}

//postメゾットのチュートリアル
func uploadMessage(c *gin.Context) {
    var req model.MessageModel
    c.BindJSON(&req)
    mess := model.MessageModel{Message: req.Message, Title: req.Title}
    c.JSON(200, mess)
}

// 一個目のAPI
func breforeResponse(c *gin.Context) {
    var req model.BeforeRequest
    c.BindJSON(&req)
    c.JSON(200, req)
}

// 二個目のAPI
func logResponse(c *gin.Context) {
    var req model.LogRequest
    c.BindJSON(&req)
    c.JSON(200, req)
}

// 三個目のAPI
func afterResponse(c *gin.Context) {
    var req model.AfterRequest
    c.BindJSON(&req)
    c.JSON(200, req)
}