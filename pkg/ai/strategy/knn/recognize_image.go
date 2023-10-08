package knn

import (
	"math/rand"
)

func (b *KnnStrategy) RecognizeImage(path string) (number int, confidence float64, _ error) {
	return int(rand.Int31n(10)), 100, nil
}
