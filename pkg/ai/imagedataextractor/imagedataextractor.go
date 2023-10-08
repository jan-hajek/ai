package imagedataextractor

import (
	"context"
	"image"
	"math"

	"github.com/jelito/ai/pkg/ai/imagex"
)

type ImageDataExtractor struct {
	xFieldsCount int
	yFieldsCount int
}

func NewImageDataExtractor(
	xFieldsCount int,
	yFieldsCount int,
) *ImageDataExtractor {
	return &ImageDataExtractor{xFieldsCount: xFieldsCount, yFieldsCount: yFieldsCount}
}

type Field struct {
	pixelCount int
	colorSum   int
}

func (i *ImageDataExtractor) ExtractFields(_ context.Context, img image.Image) []float64 {
	bounds := img.Bounds()
	width, height := bounds.Max.X, bounds.Max.Y

	fields := make([]Field, i.xFieldsCount*i.yFieldsCount)
	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			index := getQuadrantIndex(x, y, width, height, i.xFieldsCount, i.yFieldsCount)

			pixel := imagex.RgbaToPixel(img.At(x, y).RGBA())

			fields[index].pixelCount++
			fields[index].colorSum = fields[index].colorSum + pixel.R + pixel.G + pixel.B
		}
	}

	result := make([]float64, 0, i.xFieldsCount*i.yFieldsCount)
	for _, field := range fields {
		if field.pixelCount > 0 {
			colorAvg := float64(field.colorSum) / float64(field.pixelCount)
			result = append(result, colorAvg/float64(255*3))
		}
	}

	return result
}

func getQuadrantIndex(x, y, width, height, xFieldsCount, yFieldsCount int) int {
	if x >= width {
		panic("x is out of range")
	}
	if y >= height {
		panic("y is out of range")
	}
	xCount := float64(width) / float64(xFieldsCount)
	yCount := float64(height) / float64(yFieldsCount)

	xIndex := int(math.Floor(float64(x) / xCount))
	yIndex := int(math.Floor(float64(y) / yCount))

	return xIndex + yIndex*yFieldsCount
}
