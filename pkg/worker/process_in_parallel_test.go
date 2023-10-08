package worker

import (
	"context"
	"testing"

	"github.com/pkg/errors"
	"github.com/stretchr/testify/require"
)

func TestProcessInParallel(t *testing.T) {
	ctx := context.Background()
	numbers, err := ProcessInParallel(
		ctx,
		[]int{1, 2, 3, 4, 5, 6},
		func(_ context.Context, x int) (int, error) {
			return x * 2, nil
		},
		4,
	)

	require.NoError(t, err)
	require.Len(t, numbers, 6)
	require.ElementsMatch(t, []int{2, 4, 6, 8, 10, 12}, numbers)
}

func TestError(t *testing.T) {
	ctx := context.Background()
	_, err := ProcessInParallel(
		ctx,
		[]int{1, 2, 3, 4, 5, 6},
		func(_ context.Context, x int) (int, error) {
			return 0, errors.New("test error")
		},
		4,
	)

	require.ErrorContains(t, err, "test error", err)
}
