package main

import(
    "github.com/gin-gonic/gin"
    "./model"
    //"net/http"
)

func main() {
	r := gin.Default()
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
