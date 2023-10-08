package random

import (
	"context"
	"math/rand"
)

type RandomStrategy struct {
}

func New() *RandomStrategy {
	return &RandomStrategy{}
}

func (b *RandomStrategy) DataExtraction(ctx context.Context) error {
	return nil
}

func (b *RandomStrategy) PrepareSets(ctx context.Context) error {
	return nil
}

func (b *RandomStrategy) TrainAlgorithm(ctx context.Context) error {
	return nil
}

func (b *RandomStrategy) TestAlgorithm(ctx context.Context) error {
	return nil
}

func (b *RandomStrategy) RecognizeImage(path string) (number int, confidence float64, _ error) {
	return int(rand.Int31n(10)), 100, nil
}
