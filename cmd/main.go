package main

import (
	"fmt"

	"github.com/SergeyMilch/avatario/pkg/handlers"
	"github.com/SergeyMilch/avatario/pkg/service"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	// Создаем сервис и обработчик
	avatarService := *service.NewAvatarService()
	avatarHandler := handlers.NewAvatarHandler(avatarService)

	// Загрузка формы
	router.LoadHTMLGlob("templates/*")
	router.GET("/upload", avatarHandler.ShowUploadForm)

	// Обработка загруженного изображения
	router.POST("/upload", avatarHandler.Upload)

	port := ":8090"
	fmt.Printf("Сервер запущен на порту %s\n", port)
	router.Run(port)
}
