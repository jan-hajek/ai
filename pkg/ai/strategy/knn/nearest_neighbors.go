package knn

import (
	"math"
	"sort"

	"github.com/jelito/ai/pkg/ai/mathx"
)

func nearestNeighbors(k int, item Item, set Set) []Item {
	sort.Slice(set, func(i, j int) bool {
		return getDistance(item, set[i]) < getDistance(item, set[j])
	})

	return set[0:mathx.Min(k, len(set))]
}

func getDistance(validationItem Item, trainingItem Item) float64 {
	sum := 0.0
	for i := 0; i < len(validationItem.fieldsAvgs); i++ {
		sum += math.Pow(validationItem.fieldsAvgs[i]-trainingItem.fieldsAvgs[i], 2)
	}
	return sum
}
