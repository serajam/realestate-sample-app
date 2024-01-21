/* Copyright (C) Fedir Petryk */

package properties

import (
	"bytes"
	"fmt"
	"image"
	"image/jpeg"
	"math"
	"mime/multipart"

	"github.com/anthonynsimon/bild/transform"
	"github.com/google/uuid"
	"github.com/serajam/realestate-sample-app/internal/core/domain/properties"
)

type ImagesProcessService struct {
	dimensions []properties.Dimensions
	formats    map[string]struct{}
}

func NewImagesProcessService(dimensions []properties.Dimensions, formats []string) *ImagesProcessService {
	formatsMap := make(map[string]struct{}, len(formats))
	for _, format := range formats {
		formatsMap[format] = struct{}{}
	}

	return &ImagesProcessService{
		dimensions: dimensions,
		formats:    formatsMap,
	}
}

func (s ImagesProcessService) Process(file *multipart.FileHeader) (*properties.Image, error) {
	src, err := file.Open()
	if err != nil {
		return nil, err
	}
	defer src.Close()

	img, format, err := image.Decode(src)
	if err != nil {
		return nil, err
	}

	if _, ok := s.formats[format]; !ok {
		return nil, fmt.Errorf("unsupported image format: %s", format)
	}

	image := properties.Image{
		ID: uuid.New(),
	}
	for _, dimension := range s.dimensions {
		content, err := s.resize(img, file.Size, dimension.Width, dimension.Height)
		if err != nil {
			return nil, err
		}

		content.ImageType = dimension.ImageType
		image.Contents = append(image.Contents, *content)
	}

	return &image, nil
}

func (s ImagesProcessService) resize(img image.Image, size int64, width, height float64) (*properties.Content, error) {
	buf := make([]byte, 0, size)
	dst := bytes.NewBuffer(buf)

	originWidth := float64(img.Bounds().Dx())
	originHeight := float64(img.Bounds().Dy())

	if originWidth > width {
		NewHeight := math.Floor(width / originWidth * originHeight)

		// resize
		resized := transform.Resize(img, int(width), int(NewHeight), transform.Linear)
		err := jpeg.Encode(dst, resized, &jpeg.Options{Quality: 80})
		if err != nil {
			return nil, fmt.Errorf("failed to encode image: %s", err)
		}

		originWidth = float64(img.Bounds().Dx())
		originHeight = float64(img.Bounds().Dy())
	}

	if originHeight > height {
		NewWidth := math.Floor(height / originHeight * originWidth)

		// resize
		resized := transform.Resize(img, int(NewWidth), int(height), transform.Linear)
		err := jpeg.Encode(dst, resized, &jpeg.Options{Quality: 80})
		if err != nil {
			return nil, fmt.Errorf("failed to encode image: %s", err)
		}
	}

	if originHeight < height && originWidth < width {
		err := jpeg.Encode(dst, img, &jpeg.Options{Quality: 80})
		if err != nil {
			return nil, fmt.Errorf("failed to encode image: %s", err)
		}
	}

	image := properties.Content{
		Content: dst,
		Size:    int64(dst.Len()),
	}

	return &image, nil
}
