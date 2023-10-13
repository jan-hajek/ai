package knn

type results struct {
	validationDataSize int
	kResults           map[int]*kResult
}

func newResults(kList []int, validationDataSize int) *results {
	r := results{
		validationDataSize: validationDataSize,
		kResults:           make(map[int]*kResult, len(kList)),
	}
	for _, k := range kList {
		r.kResults[k] = &kResult{}
	}
	return &r
}

type kResult struct {
	triesCount        int
	correctGuessCount int
}

func (r *results) correctKGuess(k int) {
	r.kResults[k].correctGuessCount++
}

func (r *results) successRate(k int) float64 {
	return float64(r.kResults[k].correctGuessCount) / float64(r.validationDataSize)
}
