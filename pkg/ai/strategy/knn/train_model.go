package knn

import (
	"context"
	"fmt"
	"math"
	"strconv"

	"github.com/jelito/ai/pkg/ai/csvx"
	"github.com/jelito/ai/pkg/ai/mathx"
	"github.com/pkg/errors"
)

func (b *KnnStrategy) TrainModel(_ context.Context) error {
	trainData, err := csvx.ReadFromFile(b.settings.TrainingSetPath)
	if err != nil {
		return err
	}

	sets, err := splitIntoSets(trainData, b.settings.TrainModelBucketsCount)
	if err != nil {
		return err
	}

	results := newResults(b.settings.TrainModelKList, len(trainData))

	for i := 0; i < b.settings.TrainModelBucketsCount; i++ {
		validationSet, trainingSet := getTrainingAndValidationSet(i, sets)

		fmt.Printf(
			"Bucket #%d, training set: %d, validation set: %d \n",
			i+1,
			len(trainingSet),
			len(validationSet),
		)

		testBucket(results, b.settings.TrainModelKList, trainingSet, validationSet)
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

func testBucket(results *results, kList []int, trainingSet Set, validationSet Set) {
	for _, item := range validationSet {
		neighbors := nearestNeighbors(mathx.Max(kList...), item, trainingSet)

		for _, k := range kList {
			number := guessNumber(neighbors[:k])
			if number == item.number {
				results.correctKGuess(k)
			}
		}
	}
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

type Item struct {
	number         int
	sourceFileName string
	fieldsAvgs     []float64
}

type Set []Item

func splitIntoSets(trainData []csvx.Row, bucketCount int) (sets []Set, _ error) {
	length := len(trainData)
	setSize := int(math.Ceil(float64(length) / float64(bucketCount)))

	counter := 0
	set := Set{}
	for index, row := range trainData {
		number, err := strconv.Atoi(row[0])
		if err != nil {
			return nil, errors.WithStack(err)
		}

		avgs, err := convertStringsIntoFloats(row[2:])
		if err != nil {
			return nil, err
		}
		set = append(set, Item{
			number:         number,
			sourceFileName: row[1],
			fieldsAvgs:     avgs,
		})

		if index == length-1 {
			sets = append(sets, set)
			break
		}

		counter++
		if counter == setSize {
			sets = append(sets, set)
			set = Set{}
			counter = 0
		}
	}

	return
}

func convertStringsIntoFloats(input []string) ([]float64, error) {
	result := make([]float64, len(input))
	for i, value := range input {
		var err error
		result[i], err = strconv.ParseFloat(value, 64)
		if err != nil {
			return nil, errors.WithStack(err)
		}
	}
	return result, nil
}

func getTrainingAndValidationSet(i int, sets []Set) (validationSet, trainingSet Set) {
	validationSet = sets[i]

	for _, set := range sets[:i] {
		trainingSet = append(trainingSet, set...)
	}
	for _, set := range sets[i+1:] {
		trainingSet = append(trainingSet, set...)
	}

	return validationSet, trainingSet
}
