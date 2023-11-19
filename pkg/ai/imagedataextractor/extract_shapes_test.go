package imagedataextractor

import (
	"os"
	"path"
	"testing"

	"github.com/jan-hajek/ai/pkg/ai/imagex"
	"github.com/stretchr/testify/require"
)

func TestExtractShapes(t *testing.T) {
	dir, err := os.Getwd()
	require.NoError(t, err)

	img, err := imagex.OpenImage(path.Join(dir, "test_data", "test_shape.jpg"))
	require.NoError(t, err)

	ise := NewImageShapeExtractor()
	ise.ExtractShapes(nil, img)
}
