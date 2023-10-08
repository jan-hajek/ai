package imagedataextractor

import (
	"math"
	"os"
	"path"
	"testing"

	"github.com/jelito/ai/pkg/ai/imagex"
	"github.com/stretchr/testify/require"
)

func TestExtractFields(t *testing.T) {
	dir, err := os.Getwd()
	require.NoError(t, err)

	img, err := imagex.OpenImage(path.Join(dir, "test_data", "test-image.png"))
	require.NoError(t, err)

	ide := NewImageDataExtractor(3, 3)
	avgs := ide.ExtractFields(nil, img)

	require.Len(t, avgs, 9)
	require.Equal(t, 0.0, avgs[0])
	require.Equal(t, 2.0/3.0, avgs[1])
	require.Equal(t, 1.0/3.0, avgs[2])
	require.Equal(t, 2.0/3.0, avgs[3])
	require.Greater(t, math.Floor((177.0*3.0/(255.0*3.0))*10000), math.Floor(avgs[4]*10000))
	require.Greater(t, math.Floor((177.0+255.0/(255.0*3.0))*10000), math.Floor(avgs[5]*10000))
	require.Equal(t, 1.0, avgs[6])
	require.Equal(t, 5.0/9.0, avgs[7])
	require.Equal(t, 8.0/9.0, avgs[8])
}
