package knn

import (
	"context"
	"fmt"
	"runtime"
	"sync"

	"github.com/jan-hajek/ai/pkg/ai/csvx"
	"github.com/jan-hajek/ai/pkg/worker"
	"github.com/schollz/progressbar/v3"
)

type mistakeDetails struct {
	path        string
	correct     int
	guessed     int
	probGuessed float64
	secondBest  int
}

func (b *KnnStrategy) TestModel(ctx context.Context) error {
	testDataCsv, err := csvx.ReadFromFile(b.settings.TestingSetPath)
	if err != nil {
		return err
	}
	trainDataCsv, err := csvx.ReadFromFile(b.settings.TrainingSetPath)
	if err != nil {
		return err
	}

	testData, err := rowsToItems(testDataCsv)
	if err != nil {
		return err
	}

	trainData, err := rowsToItems(trainDataCsv)
	if err != nil {
		return err
	}

	k := 4
	correctGuessCount := 0
	correctPositiveGuessCount := make(map[int]int)
	correctNegativeGuessCount := make(map[int]int)
	totalCountsPerNum := make(map[int]int)
	exampleOfMistakes := make(map[int]mistakeDetails)
	maxMistakesToDisplay := 5

	curMistakesToDisplay := 0

	fmt.Printf("Testing model with k: %d, test set: %d, train set: %d \n", k, len(testData), len(trainData))

	bar := progressbar.Default(int64(len(testData)))

	var mx sync.Mutex

	worker.ProcessInParallel(
		ctx,
		testData,
		func(ctx context.Context, item Item) (interface{}, error) {
			neighbors := nearestNeighbors(k, item, trainData)

			number, numberCount, secondBest := guessNumbers(neighbors[:k])

			func() {
				mx.Lock()
				defer mx.Unlock()

				totalCountsPerNum[item.number]++

				if number == item.number {
					correctGuessCount++
					correctPositiveGuessCount[number]++
				} else {
					for i := 0; i < 10; i++ {
						if i != number {
							correctNegativeGuessCount[number]++
						}
					}
					// create mistake example
					if curMistakesToDisplay < maxMistakesToDisplay {
						exampleOfMistakes[curMistakesToDisplay] = mistakeDetails{
							path:        item.sourceFileName,
							correct:     item.number,
							guessed:     number,
							probGuessed: float64(numberCount) / float64(k),
							secondBest:  secondBest,
						}
						curMistakesToDisplay++
					}
				}
			}()

			bar.Add(1)

			return nil, nil
		},
		worker.WithWorkersCount(runtime.NumCPU()),
	)

	bar.Exit()

	fmt.Printf(
		"Accuracy: k=%d, accuracy=%.2f%%\n",
		k,
		float64(correctGuessCount)/float64(len(testDataCsv))*100,
	)
	for i := 0; i < 10; i++ {
		fmt.Printf(
			"Num=%d: TotalCount=%d, Sensitivity=%.2f%%, Specificity=%.2f%%\n",
			i,
			totalCountsPerNum[i],
			float64(correctPositiveGuessCount[i])/float64(totalCountsPerNum[i])*100,
			float64(correctNegativeGuessCount[i])/float64((len(testDataCsv))-totalCountsPerNum[i])*100,
		)
	}
	for i := 0; i < maxMistakesToDisplay; i++ {
		fmt.Printf(
			"Mistakes examples:\n %s, Correct=%d, Guessed=%d, ProbGuessed=%.2f%%, SecondBest=%d\n",
			exampleOfMistakes[i].path,
			exampleOfMistakes[i].correct,
			exampleOfMistakes[i].guessed,
			exampleOfMistakes[i].probGuessed*100,
			exampleOfMistakes[i].secondBest,
		)
	}
	return nil
}

func guessNumbers(nearestNeighbors []Item) (int, int, int) {
	mostCommon := make(map[int]int)
	bestNumber := -1
	bestCount := 0
	secondBest := -1
	for _, n := range nearestNeighbors {
		count := mostCommon[n.number] + 1
		mostCommon[n.number] = count
		if count > bestCount {
			bestCount = count
			secondBest = bestNumber
			bestNumber = n.number
		}
	}

	return bestNumber, bestCount, secondBest
}
