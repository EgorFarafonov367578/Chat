package main

import (
	"chat/internal/domain"
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

var chatMessages = []domain.Message{
	{
		Text: "message1",
		Time: time.Now()},
	{
		Text: "message2",
		Time: time.Now()},
	{
		Text: "message3",
		Time: time.Now()},
}

func main() {
	r := gin.Default()
	r.GET("/", func(ctx *gin.Context) {
		ctx.File("index.html")
		fmt.Println("Done")
	})
	r.GET("/messages", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, chatMessages)
	})
	_ = r.Run()
	fmt.Println("Stop")
}
