package main

import(
    "github.com/gin-gonic/gin"
    "github.com/jinzhu/gorm"
    "./model"
    _ "github.com/mattn/go-sqlite3"
    //"net/http"
)

func main() {
	r := gin.Default()

    dbInit()
    dbInsert("sample", "stats")
    
	r.GET("/ping", helloWorld)
    r.POST("/post", uploadMessage)
	r.Run()
}

func helloWorld(c *gin.Context) {
    c.JSON(200, gin.H{
        "message": "helloworld",
    })
}

func uploadMessage(c *gin.Context) {
    var req model.MessageModel
    c.BindJSON(&req)
    mess := model.MessageModel{Message: req.Message, Title: req.Title}
    c.JSON(200, mess)
}

//DB初期化
func dbInit() {
    db, err := gorm.Open("sqlite3", "model/database.sqlite3")
    if err != nil {
        panic("データベース開けず！（dbInit）")
    }
    db.AutoMigrate(&model.BeforeModel{})
    db.AutoMigrate(&model.LogModel{})
    defer db.Close()
}

//CREATE
func dbInsert(text string, status string) {
    db, err := gorm.Open("sqlite3", "model/database.sqlite3")
    if err != nil {
        panic("データベース開けず！（dbInsert)")
    }
    db.Create(&model.BeforeModel{Text: text, Status: status})
    defer db.Close()
}