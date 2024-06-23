package main

import (
	"chat/internal/domain"
	mr "chat/internal/repositories/message/inmemory"
	ur "chat/internal/repositories/user/inmemory"
	"chat/internal/usecase"
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

/*var chatMessages = []domain.Message{
	{
		Text: "message1",
		Time: time.Now()},
	{
		Text: "message2",
		Time: time.Now()},
	{
		Text: "message3",
		Time: time.Now()},
}*/

func main() {
	mr := mr.NewMessageRepository()
	ur := ur.NewUserRepository()
	messageService := usecase.NewMessage(mr, ur)
	userService := usecase.NewUser(ur)
	userService.RegisterUser(&domain.User{
		Name: "admin",
	})
	messageService.RegisterMessage(&domain.Message{
		Text:       "Hello",
		Time:       time.Now(),
		FromUserId: 0,
	})
	messageService.RegisterMessage(&domain.Message{
		Text:       "world",
		Time:       time.Now(),
		FromUserId: 0,
	})
	r := gin.Default()
	r.GET("/", func(ctx *gin.Context) {
		ctx.File("index.html")
		fmt.Println("Done")
	})
	r.GET("/messages", func(ctx *gin.Context) {
		messages, _ := messageService.GetMessages()
		ctx.JSON(http.StatusOK, messages)
	})
	_ = r.Run()
	fmt.Println("Stop")
}
