package service

import (
	"bytes"
	"image"
	"image/color"
	"image/draw"
	"image/jpeg"
	"image/png"
	"io"

	"github.com/nfnt/resize"
)

type AvatarService struct{}

func NewAvatarService() *AvatarService {
	return &AvatarService{}
}

func (as *AvatarService) ProcessImage(file io.Reader, shape string) ([]byte, error) {
	// Преобразуем изображение
	img, _, err := image.Decode(file)
	if err != nil {
		return nil, err
	}

	// Уменьшаем изображение с использованием библиотеки nfnt/resize
	newSize := 150
	img = resize.Resize(uint(newSize), uint(newSize), img, resize.Lanczos3)

	// Преобразуем в формат RGBA, чтобы иметь возможность управлять альфа-каналом
	rgba := image.NewRGBA(img.Bounds())
	draw.Draw(rgba, rgba.Bounds(), img, image.Point{}, draw.Src)

	// Создаем маску круга, если форма "circle"
	if shape == "circle" {
		center := image.Point{newSize / 2, newSize / 2}
		radius := newSize / 2

		for y := 0; y < newSize; y++ {
			for x := 0; x < newSize; x++ {
				dx := x - center.X
				dy := y - center.Y
				distance := dx*dx + dy*dy

				if distance > radius*radius {
					// Если пиксель находится за пределами круга, делаем его прозрачным
					rgba.Set(x, y, color.RGBA{0, 0, 0, 0})
				}
			}
		}
	}

	// Добавляем рамку вокруг круглой области аватарки
	borderColor := color.RGBA{169, 169, 169, 255} // Серый цвет
	borderWidth := 1

	center := image.Point{newSize / 2, newSize / 2}
	radius := newSize / 2

	for y := 0; y < newSize; y++ {
		for x := 0; x < newSize; x++ {
			dx := x - center.X
			dy := y - center.Y
			distance := dx*dx + dy*dy

			if distance <= radius*radius && distance > (radius-borderWidth)*(radius-borderWidth) {
				rgba.Set(x, y, borderColor)
			}
		}
	}

	// Кодируем изображение обратно в формат JPEG или PNG в зависимости от исходного формата
	var buf bytes.Buffer
	if shape == "circle" {
		err = png.Encode(&buf, rgba)
	} else {
		err = jpeg.Encode(&buf, rgba, nil)
	}

	if err != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}
