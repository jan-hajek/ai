package knn

import (
	"math"
	"sort"

	"github.com/jan-hajek/ai/pkg/ai/mathx"
)

func nearestNeighbors(k int, item Item, allItems []Item) []Item {
	allItemsCopied := make([]Item, len(allItems))
	copy(allItemsCopied, allItems)
	sort.Slice(allItemsCopied, func(i, j int) bool {
		return getDistance(item, allItemsCopied[i]) < getDistance(item, allItemsCopied[j])
	})

	return allItemsCopied[0:mathx.Min(k, len(allItemsCopied))]
}

func getDistance(validationItem Item, trainingItem Item) float64 {
	sum := 0.0
	for i := 0; i < len(validationItem.fieldsAvgs); i++ {
		sum += math.Pow(validationItem.fieldsAvgs[i]-trainingItem.fieldsAvgs[i], 2)
	}
	return sum
}
