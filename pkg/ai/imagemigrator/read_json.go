package imagemigrator

import (
	"context"
	"encoding/json"
	"io"
	"os"
	"path"
	"time"
)

func (i *ImageMigrator) readJson(ctx context.Context) ([]extractImagesInput, error) {
	jsonFile, err := os.Open(path.Join(i.origDataDir, i.annotationFileName))
	if err != nil {
		return nil, err
	}
	defer jsonFile.Close()

	byteValue, _ := io.ReadAll(jsonFile)

	var annotation Annotation

	err = json.Unmarshal(byteValue, &annotation)
	if err != nil {
		return nil, err
	}

	result := []extractImagesInput{}

	for _, image := range annotation.Images {
		result = append(result, extractImagesInput{
			imagePath: image.FileName,
			coords: []coords{
				{
					x:      0,
					y:      0,
					width:  0,
					height: 0,
					number: 0,
				},
			},
		})
	}

	return result, nil
}

type Annotation struct {
	Info struct {
		Year        string    `json:"year"`
		Version     string    `json:"version"`
		Description string    `json:"description"`
		Contributor string    `json:"contributor"`
		Url         string    `json:"url"`
		DateCreated time.Time `json:"date_created"`
	} `json:"info"`
	Licenses []struct {
		Id   int    `json:"id"`
		Url  string `json:"url"`
		Name string `json:"name"`
	} `json:"licenses"`
	Categories []struct {
		Id            int    `json:"id"`
		Name          string `json:"name"`
		Supercategory string `json:"supercategory"`
	} `json:"categories"`
	Images []struct {
		Id           int       `json:"id"`
		License      int       `json:"license"`
		FileName     string    `json:"file_name"`
		Height       int       `json:"height"`
		Width        int       `json:"width"`
		DateCaptured time.Time `json:"date_captured"`
	} `json:"images"`
	Annotations []struct {
		Id           int           `json:"id"`
		ImageId      int           `json:"image_id"`
		CategoryId   int           `json:"category_id"`
		Bbox         []int         `json:"bbox"`
		Area         int           `json:"area"`
		Segmentation []interface{} `json:"segmentation"`
		Iscrowd      int           `json:"iscrowd"`
	} `json:"annotations"`
}
