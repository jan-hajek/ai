package knn

import (
	"context"
)

func (b *KnnStrategy) TrainModel(ctx context.Context) error {
	// split training data into x buckets
	// one bucket is validation set, others are training sets
	// for each item in the validation set find k nearest neighbors in the training set
	// calculate accuracy
	// repeat for each bucket
	// average accuracy
	// repeat for each k
	// pick the best k
	// save the model
	return nil
}
