package handlers

import (
	"net/http"

	"github.com/SergeyMilch/avatario/pkg/service"
	"github.com/gin-gonic/gin"
)

type AvatarHandler struct {
	avatarService service.AvatarService
}

func NewAvatarHandler(as service.AvatarService) *AvatarHandler {
	return &AvatarHandler{
		avatarService: as,
	}
}

func (ah *AvatarHandler) ShowUploadForm(c *gin.Context) {
	c.HTML(http.StatusOK, "upload.html", nil)
}

func (ah *AvatarHandler) Upload(c *gin.Context) {
	shape := c.PostForm("shape") // Выбор формы аватарки (круглая или квадратная)

	file, err := c.FormFile("file")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	fileData, err := file.Open()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	processedImage, err := ah.avatarService.ProcessImage(fileData, shape)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.Header("Content-Disposition", "attachment; filename=processed_image.jpg")
	c.Data(http.StatusOK, "image/jpeg", processedImage)
}
