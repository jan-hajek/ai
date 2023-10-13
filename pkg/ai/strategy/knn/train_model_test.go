package knn

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestSplitIntoSets1(t *testing.T) {
	sets, err := splitIntoSets([][]string{
		{"1", "file1", "0.1", "0.2"},
		{"2", "file2", "0.3", "0.4"},
		{"3", "file3", "0.5", "0.6"},
		{"4", "file4", "0.7", "0.8"},
		{"5", "file5", "0.9", "1.0"},
		{"6", "file6", "1.1", "1.2"},
		{"7", "file7", "1.3", "1.4"},
	}, 4)
	require.NoError(t, err)
	require.Equal(t, 4, len(sets))
	require.Equal(t, sets[0], Set{
		{
			number:         1,
			sourceFileName: "file1",
			fieldsAvgs:     []float64{0.1, 0.2},
		},
		{
			number:         2,
			sourceFileName: "file2",
			fieldsAvgs:     []float64{0.3, 0.4},
		},
	})

	require.Equal(t, sets[3], Set{
		{
			number:         7,
			sourceFileName: "file7",
			fieldsAvgs:     []float64{1.3, 1.4},
		},
	})
}

func TestSplitIntoSets2(t *testing.T) {
	sets, err := splitIntoSets([][]string{
		{"1", "file1", "0.1", "0.2"},
		{"2", "file2", "0.3", "0.4"},
		{"3", "file3", "0.5", "0.6"},
		{"4", "file4", "0.7", "0.8"},
		{"5", "file5", "0.9", "1.0"},
		{"6", "file6", "1.1", "1.2"},
		{"7", "file7", "1.3", "1.4"},
		{"8", "file8", "1.5", "1.6"},
	}, 4)
	require.NoError(t, err)
	require.Equal(t, sets[0], Set{
		{
			number:         1,
			sourceFileName: "file1",
			fieldsAvgs:     []float64{0.1, 0.2},
		},
		{
			number:         2,
			sourceFileName: "file2",
			fieldsAvgs:     []float64{0.3, 0.4},
		},
	})

	require.Equal(t, sets[3], Set{
		{
			number:         7,
			sourceFileName: "file7",
			fieldsAvgs:     []float64{1.3, 1.4},
		},
		{
			number:         8,
			sourceFileName: "file8",
			fieldsAvgs:     []float64{1.5, 1.6},
		},
	})
}

func TestSplitIntoSets3(t *testing.T) {
	sets := []Set{
		{{number: 1}},
		{{number: 2}},
		{{number: 3}},
		{{number: 4}},
	}
	{
		validationSet, trainingSet := getTrainingAndValidationSet(0, sets)
		require.Equal(t, Set{{number: 1}}, validationSet)
		require.Equal(t, Set{{number: 2}, {number: 3}, {number: 4}}, trainingSet)
	}
	{
		validationSet, trainingSet := getTrainingAndValidationSet(1, sets)
		require.Equal(t, Set{{number: 2}}, validationSet)
		require.Equal(t, Set{{number: 1}, {number: 3}, {number: 4}}, trainingSet)
	}
	{
		validationSet, trainingSet := getTrainingAndValidationSet(2, sets)
		require.Equal(t, Set{{number: 3}}, validationSet)
		require.Equal(t, Set{{number: 1}, {number: 2}, {number: 4}}, trainingSet)
	}
	{
		validationSet, trainingSet := getTrainingAndValidationSet(3, sets)
		require.Equal(t, Set{{number: 4}}, validationSet)
		require.Equal(t, Set{{number: 1}, {number: 2}, {number: 3}}, trainingSet)
	}
}
