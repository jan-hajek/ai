package imagedataextractor

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestGetQuadrantIndex(t *testing.T) {
	require.Equal(t, 0, getQuadrantIndex(0, 0, 100, 100, 3, 3))
	require.Equal(t, 0, getQuadrantIndex(0, 33, 100, 100, 3, 3))
	require.Equal(t, 0, getQuadrantIndex(33, 0, 100, 100, 3, 3))
	require.Equal(t, 0, getQuadrantIndex(33, 33, 100, 100, 3, 3))

	require.Equal(t, 1, getQuadrantIndex(34, 0, 100, 100, 3, 3))
	require.Equal(t, 1, getQuadrantIndex(34, 33, 100, 100, 3, 3))
	require.Equal(t, 1, getQuadrantIndex(66, 0, 100, 100, 3, 3))
	require.Equal(t, 1, getQuadrantIndex(66, 33, 100, 100, 3, 3))

	require.Equal(t, 2, getQuadrantIndex(67, 0, 100, 100, 3, 3))
	require.Equal(t, 2, getQuadrantIndex(67, 33, 100, 100, 3, 3))
	require.Equal(t, 2, getQuadrantIndex(99, 0, 100, 100, 3, 3))
	require.Equal(t, 2, getQuadrantIndex(99, 33, 100, 100, 3, 3))

	require.Equal(t, 3, getQuadrantIndex(0, 34, 100, 100, 3, 3))
	require.Equal(t, 3, getQuadrantIndex(0, 66, 100, 100, 3, 3))
	require.Equal(t, 3, getQuadrantIndex(33, 34, 100, 100, 3, 3))
	require.Equal(t, 3, getQuadrantIndex(33, 66, 100, 100, 3, 3))

	require.Equal(t, 4, getQuadrantIndex(34, 34, 100, 100, 3, 3))
	require.Equal(t, 4, getQuadrantIndex(34, 66, 100, 100, 3, 3))
	require.Equal(t, 4, getQuadrantIndex(66, 34, 100, 100, 3, 3))
	require.Equal(t, 4, getQuadrantIndex(66, 66, 100, 100, 3, 3))

	require.Equal(t, 5, getQuadrantIndex(67, 34, 100, 100, 3, 3))
	require.Equal(t, 5, getQuadrantIndex(67, 66, 100, 100, 3, 3))
	require.Equal(t, 5, getQuadrantIndex(99, 34, 100, 100, 3, 3))
	require.Equal(t, 5, getQuadrantIndex(99, 66, 100, 100, 3, 3))

	require.Equal(t, 6, getQuadrantIndex(0, 67, 100, 100, 3, 3))
	require.Equal(t, 6, getQuadrantIndex(0, 99, 100, 100, 3, 3))
	require.Equal(t, 6, getQuadrantIndex(33, 67, 100, 100, 3, 3))
	require.Equal(t, 6, getQuadrantIndex(33, 99, 100, 100, 3, 3))

	require.Equal(t, 7, getQuadrantIndex(34, 67, 100, 100, 3, 3))
	require.Equal(t, 7, getQuadrantIndex(34, 99, 100, 100, 3, 3))
	require.Equal(t, 7, getQuadrantIndex(66, 67, 100, 100, 3, 3))
	require.Equal(t, 7, getQuadrantIndex(66, 99, 100, 100, 3, 3))

	require.Equal(t, 8, getQuadrantIndex(67, 67, 100, 100, 3, 3))
	require.Equal(t, 8, getQuadrantIndex(67, 99, 100, 100, 3, 3))
	require.Equal(t, 8, getQuadrantIndex(99, 67, 100, 100, 3, 3))
	require.Equal(t, 8, getQuadrantIndex(99, 99, 100, 100, 3, 3))
}
