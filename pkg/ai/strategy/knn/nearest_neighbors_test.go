package knn

import (
	"strconv"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestNearestNeighbors(t *testing.T) {
	testCases := []struct {
		k      int
		item   Item
		set    Set
		result []Item
	}{
		{
			k:    3,
			item: testingItem(10, 10),
			set: []Item{
				testingItem(20, 20), // 10
				testingItem(0, 0),   // 10
				testingItem(19, 19), // 9
				testingItem(1, 1),   // 9
				testingItem(18, 18), // 8
				testingItem(2, 2),   // 8
				testingItem(17, 17), // 7
				testingItem(3, 3),   // 7
			},
			result: []Item{
				testingItem(17, 17), //9
				testingItem(3, 3),   // 10
				testingItem(18, 18), // 9
			},
		},
		{
			k:    3,
			item: testingItem(10, 10),
			set: []Item{
				testingItem(0, 0),   // 10
				testingItem(18, 18), // 8
				testingItem(20, 20), // 10
				testingItem(17, 17), // 7
				testingItem(2, 2),   // 8
				testingItem(3, 3),   // 7
				testingItem(19, 19), // 9
				testingItem(1, 1),   // 9
			},
			result: []Item{
				testingItem(17, 17), //9
				testingItem(3, 3),   // 10
				testingItem(18, 18), // 9
			},
		},
	}

	for index, testCase := range testCases {
		t.Run(strconv.Itoa(index), func(t *testing.T) {
			require.Equal(t, testCase.result, nearestNeighbors(
				testCase.k,
				testCase.item,
				testCase.set,
			))
		})
	}
}

func TestGetDistance(t *testing.T) {
	testCases := []struct {
		x      []float64
		y      []float64
		result float64
	}{
		{
			x:      []float64{1, 1},
			y:      []float64{2, 2},
			result: 2,
		},
		{
			x:      []float64{2, 2},
			y:      []float64{1, 1},
			result: 2,
		},
		{
			x:      []float64{3, 1},
			y:      []float64{1, 3},
			result: 8,
		},
		{
			x:      []float64{1, 3, 2},
			y:      []float64{2, 1, 5},
			result: 1 + 4 + 9,
		},
	}

	for index, testCase := range testCases {
		t.Run(strconv.Itoa(index), func(t *testing.T) {
			require.Equal(t, testCase.result, getDistance(testingItem(testCase.x...), testingItem(testCase.y...)))
		})
	}
}

func testingItem(avgs ...float64) Item {
	return Item{
		fieldsAvgs: avgs,
	}
}
