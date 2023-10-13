package knn

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestGetTrainingAndValidationSet(t *testing.T) {
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
