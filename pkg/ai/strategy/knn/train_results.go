package knn

import "sync"

type prediction struct {
	item Item
	predicted int
	k int
}

type results struct {
	validationDataSize int
	kResults           map[int]*kResult
	mx                 sync.Mutex
	falsePredictions []prediction
}

func newResults(kList []int, validationDataSize int) *results {
	r := results{
		validationDataSize: validationDataSize,
		kResults:           make(map[int]*kResult, len(kList)),
		falsePredictions: []prediction{},
	}
	for _, k := range kList {
		r.kResults[k] = &kResult{}
	}
	return &r
}

type kResult struct {
	triesCount        int
	correctGuessCount int
	correct map[int]*float64
	// probability that number correctly wasn't guessed
	specificity map[int]*float64
}

func (r *results) correctKGuess(k int) {
	r.mx.Lock()
	defer r.mx.Unlock()
	r.kResults[k].correctGuessCount++
}

func (r *results) successRate(k int) float64 {
	return float64(r.kResults[k].correctGuessCount) / float64(r.validationDataSize)
}

func (r *results) addFalsePrediction(item Item, predicted int, k int) {
	r.falsePredictions = append(r.falsePredictions, prediction{item: item, predicted: predicted, k: k})
}

