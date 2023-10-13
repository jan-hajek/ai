package knn

import (
	"context"
	"fmt"
	"runtime"

	"github.com/jan-hajek/ai/pkg/ai/csvx"
	"github.com/jan-hajek/ai/pkg/ai/mathx"
	"github.com/jan-hajek/ai/pkg/ai/slices"
	"github.com/jan-hajek/ai/pkg/worker"
	"github.com/schollz/progressbar/v3"
)

func (b *KnnStrategy) TrainModel(ctx context.Context) error {
	trainData, err := csvx.ReadFromFile(b.settings.TrainingSetPath)
	if err != nil {
		return err
	}

	allData, err := rowsToItems(trainData)
	if err != nil {
		return err
	}

	sets := slices.SplitByBucketCount(allData, b.settings.TrainModelBucketsCount)

	results := newResults(b.settings.TrainModelKList, len(trainData))

	for i := 0; i < b.settings.TrainModelBucketsCount; i++ {
		validationSet, trainingSet := getTrainingAndValidationSet(i, sets)

		fmt.Printf(
			"Bucket #%d, training set: %d, validation set: %d \n",
			i+1,
			len(trainingSet),
			len(validationSet),
		)

		testBucket(ctx, results, b.settings.TrainModelKList, validationSet, trainingSet)
	}

	bestAccuracy := 0.0
	bestAccuracyK := 0
	for _, k := range b.settings.TrainModelKList {
		accuracy := results.successRate(k)
		fmt.Printf("k=%d, accuracy=%.2f%%\n", k, accuracy*100)
		if accuracy > bestAccuracy {
			bestAccuracy = accuracy
			bestAccuracyK = k
		}
	}

	fmt.Printf("Best accuracy: k=%d, accuracy=%.2f%%\n", bestAccuracyK, bestAccuracy*100)

	return nil
}

func testBucket(ctx context.Context, results *results, kList []int, validationSet []Item, trainingSet []Item) {
	bar := progressbar.Default(int64(len(validationSet)))

	maxK := mathx.Max(kList...)

	worker.ProcessInParallel(
		ctx,
		validationSet,
		func(ctx context.Context, item Item) (interface{}, error) {
			neighbors := nearestNeighbors(maxK, item, trainingSet)

			for _, k := range kList {
				number := guessNumber(neighbors[:k])
				if number == item.number {
					results.correctKGuess(k)
				}
			}

			bar.Add(1)

			return nil, nil
		},
		worker.WithWorkersCount(runtime.NumCPU()),
	)

	bar.Exit()
}

func guessNumber(nearestNeighbors []Item) int {
	mostCommon := make(map[int]int)
	bestNumber := -1
	bestCount := 0
	for _, n := range nearestNeighbors {
		count := mostCommon[n.number] + 1
		mostCommon[n.number] = count
		if count > bestCount {
			bestCount = count
			bestNumber = n.number
		}
	}

	return bestNumber
}

func getTrainingAndValidationSet(i int, sets [][]Item) (validationSet, trainingSet []Item) {
	validationSet = sets[i]

	for _, set := range sets[:i] {
		trainingSet = append(trainingSet, set...)
	}
	for _, set := range sets[i+1:] {
		trainingSet = append(trainingSet, set...)
	}

	return validationSet, trainingSet
}
