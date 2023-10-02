package repository

import "fmt"

type AvatarRepository struct {
	// Тут можно добавить зависимости, если потребуется
}

func NewAvatarRepository() *AvatarRepository {
	return &AvatarRepository{}
}

func (ar *AvatarRepository) SaveImage() {
	fmt.Println("Сохранение изображения")
}
