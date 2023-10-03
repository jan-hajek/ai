package imagemigrator

import (
	"os"
	"path"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestExtractImage(t *testing.T) {
	dir, err := os.Getwd()
	require.NoError(t, err)

	origDataDir := path.Join(dir, "extract_images_test_data")
	migratedDataDir := path.Join(dir, "extract_images_test_data")
	annotationFileName := ""

	input := extractImagesInput{
		imageName: "936.jpg",
		coords: []coords{
			{x: 140, y: 27, width: 22, height: 31, number: 9},
			{x: 216, y: 23, width: 22, height: 35, number: 3},
			{x: 286, y: 20, width: 19, height: 38, number: 6},
		},
	}

	m := NewImageMigrator(annotationFileName, origDataDir, migratedDataDir)
	files, err := m.extractImages(nil, input)
	require.NoError(t, err)
	require.Equal(t, 3, len(files))

	err = removeFiles(files)
	require.NoError(t, err)
}

func removeFiles(files []string) error {
	for _, file := range files {
		err := os.Remove(file)
		if err != nil {
			return err
		}
	}
	return nil
}
