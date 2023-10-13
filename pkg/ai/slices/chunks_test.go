package slices_test

import (
	"testing"

	"github.com/jan-hajek/ai/pkg/ai/slices"
	"github.com/stretchr/testify/require"
)

func TestChunks(t *testing.T) {
	actual := slices.SplitByBucketSize([]int{1, 2, 3, 4, 5, 6, 7, 8, 9}, 3)
	expected := [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}
	require.Equal(t, expected, actual)

	actual = slices.SplitByBucketSize([]int{9, 8, 7, 6, 5, 4, 3, 2, 1}, 3)
	expected = [][]int{{9, 8, 7}, {6, 5, 4}, {3, 2, 1}}
	require.Equal(t, expected, actual)

	actual = slices.SplitByBucketSize([]int{1, 2, 3, 4, 5, 6, 7, 8, 9}, 4)
	expected = [][]int{{1, 2, 3, 4}, {5, 6, 7, 8}, {9}}
	require.Equal(t, expected, actual)

	actual = slices.SplitByBucketSize([]int{1, 2, 3, 4, 5, 6, 7, 8, 9}, 2)
	expected = [][]int{{1, 2}, {3, 4}, {5, 6}, {7, 8}, {9}}
	require.Equal(t, expected, actual)

	actual = slices.SplitByBucketSize([]int{1, 2, 3, 4, 5, 6, 7, 8, 9}, 1)
	expected = [][]int{{1}, {2}, {3}, {4}, {5}, {6}, {7}, {8}, {9}}
	require.Equal(t, expected, actual)

	actual = slices.SplitByBucketSize([]int{1, 2, 3, 4, 5, 6, 7, 8, 9}, 9)
	expected = [][]int{{1, 2, 3, 4, 5, 6, 7, 8, 9}}
	require.Equal(t, expected, actual)

	actual = slices.SplitByBucketSize([]int{1, 2, 3, 4, 5, 6, 7, 8, 9}, 10)
	expected = [][]int{{1, 2, 3, 4, 5, 6, 7, 8, 9}}
	require.Equal(t, expected, actual)

	actual = slices.SplitByBucketSize([]int{1, 2, 3, 4, 5, 6, 7, 8, 9}, 0)
	expected = [][]int{{1, 2, 3, 4, 5, 6, 7, 8, 9}}
	require.Equal(t, expected, actual)

	actual = slices.SplitByBucketSize([]int{1, 2, 3, 4, 5, 6, 7, 8, 9}, -1)
	expected = [][]int{{1, 2, 3, 4, 5, 6, 7, 8, 9}}
	require.Equal(t, expected, actual)

	actual = slices.SplitByBucketSize([]int{}, 1)
	expected = [][]int{}
	require.Equal(t, expected, actual)

	actual = slices.SplitByBucketSize([]int{}, -1)
	expected = [][]int{}
	require.Equal(t, expected, actual)
}

func TestSplitByBucketCount1(t *testing.T) {
	buckets := slices.SplitByBucketCount([][]string{
		{"1", "file1", "0.1", "0.2"},
		{"2", "file2", "0.3", "0.4"},
		{"3", "file3", "0.5", "0.6"},
		{"4", "file4", "0.7", "0.8"},
		{"5", "file5", "0.9", "1.0"},
		{"6", "file6", "1.1", "1.2"},
		{"7", "file7", "1.3", "1.4"},
	}, 4)
	require.Equal(t, 4, len(buckets))
	require.Equal(t, buckets[0], [][]string{
		{"1", "file1", "0.1", "0.2"},
		{"2", "file2", "0.3", "0.4"},
	})

	require.Equal(t, buckets[3], [][]string{
		{"7", "file7", "1.3", "1.4"},
	})
}

func TestSplitByBucketCount2(t *testing.T) {
	buckets := slices.SplitByBucketCount([][]string{
		{"1", "file1", "0.1", "0.2"},
		{"2", "file2", "0.3", "0.4"},
		{"3", "file3", "0.5", "0.6"},
		{"4", "file4", "0.7", "0.8"},
		{"5", "file5", "0.9", "1.0"},
		{"6", "file6", "1.1", "1.2"},
		{"7", "file7", "1.3", "1.4"},
		{"8", "file8", "1.5", "1.6"},
	}, 4)
	require.Equal(t, buckets[0], [][]string{
		{"1", "file1", "0.1", "0.2"},
		{"2", "file2", "0.3", "0.4"},
	})

	require.Equal(t, buckets[3], [][]string{
		{"7", "file7", "1.3", "1.4"},
		{"8", "file8", "1.5", "1.6"},
	})
}
