package imagedataextractor

import (
	"context"
	"math"
)

type ImageDataExtractor struct {
	xFieldsCount int
	yFieldsCount int
}

type extractFieldsInput struct {
	pixels []Pixel
	length int
	width  int
}

type Pixel struct {
	x int
	y int
	R int
	G int
	B int
	A int
}

type Field struct {
	x          int
	y          int
	pixelCount int
	colorSum   int
	colorRatio int
}

func (i *ImageDataExtractor) ExtractFields(ctx context.Context, input extractFieldsInput) ([][]Field, error) {
	xBand := float64(input.width / i.xFieldsCount)
	yBand := float64(input.length / i.yFieldsCount)

	xDivs := []int{}
	yDivs := []int{}

	fields := [][]Field{}

	for j := 1; j <= i.xFieldsCount; j++ {
		xDivs[j-1] = int(math.Round(float64(j) * xBand))
		for k := 1; k <= i.yFieldsCount; k++ {
			yDivs[k-1] = int(math.Round(float64(k) * yBand))
			fields[j][k] = Field{j, k, 0, 0, 0}
		}
	}

	for _, pixel := range i.pixels {
		xField := 0
		for ix, xDiv := range xDivs {
			if pixel.x < xDiv {
				xField = ix + 1
			}
		}

		yField := 0
		for ix, yDiv := range yDivs {
			if pixel.y < yDiv {
				yField = ix + 1
			}
		}

		fields[xField-1][yField-1].pixelCount++
		fields[xField-1][yField-1].colorSum = fields[xField-1][yField-1].colorSum + pixel.r + pixel.g + pixel.b
	}

	for j := 1; j <= i.xFieldsCount; j++ {
		for k := 1; k <= i.yFieldsCount; k++ {
			colorAvg := fields[j][k].colorSum/fields[j][k].pixelCount
			fields[j][k].colorRatio = colorAvg/(255*3)
		}
	}

	return fields, nil
}
