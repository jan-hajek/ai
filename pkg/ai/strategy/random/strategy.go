package random

import (
	"context"
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

func (b *RandomStrategy) TrainModel(ctx context.Context) error {
	return nil
}

func (b *RandomStrategy) TestModel(ctx context.Context) error {
	return nil
}
