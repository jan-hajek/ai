package knn

import (
	"context"
	"image"
	_ "image/jpeg"
	_ "image/png"
	"sync"

	"github.com/jelito/ai/pkg/ai/imagex"
	"github.com/jelito/ai/pkg/ai/strategy"
)

func (b *Strategy) TrainFiles(ctx context.Context, files []strategy.TrainFile) error {
	ch := make(chan *strategy.TrainFile)
	wg := sync.WaitGroup{}
	workersCount := 4
	for i := 0; i < workersCount; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for file := range ch {
				q, err := b.getQuadrants(ctx, file)
				if err != nil {
					// FIXME - janhajek dodelat
					return
				}

				_ = q
			}
		}()
	}

	for index := range files {
		ch <- &files[index]
	}

	close(ch)

	wg.Wait()

	return nil
}

func (b *Strategy) getQuadrants(ctx context.Context, file *strategy.TrainFile) ([]int, error) {
	img, err := imagex.OpenImage(file.Path)
	if err != nil {
		return nil, err
	}

	_ = b.splitIntoParts(img)

	// calculate avg of each

	return []int{1, 2, 3, 4}, nil
}

func (b *Strategy) splitIntoParts(img image.Image) map[int][]imagex.Pixel {
	bounds := img.Bounds()
	width, height := bounds.Max.X, bounds.Max.Y

	widthBoundary := width / 2
	heightBoundary := height / 2

	pixels := make(map[int][]imagex.Pixel)
	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			index := 0
			switch true {
			case y < widthBoundary && x < heightBoundary:
				index = 0
			case y < widthBoundary && x >= heightBoundary:
				index = 1
			case y >= widthBoundary && x < heightBoundary:
				index = 2
			case y >= widthBoundary && x >= heightBoundary:
				index = 3
			}

			pixels[index] = append(pixels[index], imagex.RgbaToPixel(img.At(x, y).RGBA()))
		}
	}

	return pixels
}
