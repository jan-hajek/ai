package slices

import (
	"math"
)

func SplitByBucketSize[T comparable](items []T, chunkSize int) [][]T {
	if len(items) < 1 {
		return [][]T{}
	}
	if chunkSize < 1 {
		return [][]T{items}
	}

	var chunks = make([][]T, 0, (len(items)/chunkSize)+1)
	for chunkSize < len(items) {
		items, chunks = items[chunkSize:], append(chunks, items[0:chunkSize:chunkSize])
	}
	return append(chunks, items)
}

func SplitByBucketCount[T any](items []T, bucketCount int) (result [][]T) {
	length := len(items)
	setSize := int(math.Ceil(float64(length) / float64(bucketCount)))

	counter := 0
	bucket := make([]T, 0, setSize)
	for index, item := range items {
		bucket = append(bucket, item)

		if index == length-1 {
			result = append(result, bucket)
			break
		}

		counter++
		if counter == setSize {
			result = append(result, bucket)
			bucket = make([]T, 0, setSize)
			counter = 0
		}
	}

	return
}
