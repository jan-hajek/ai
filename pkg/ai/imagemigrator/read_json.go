package imagemigrator

import (
	"context"
	"encoding/json"
	"io/ioutil"
	"os"
	"path"
	"strconv"
)

func (i *ImageMigrator) readJson(ctx context.Context) ([]extractImagesInput, error) {
	annotationFile, err := os.Open(path.Join(i.sourceDir, i.annotationFileName))
	if err != nil {
		return nil, err
	}
	defer annotationFile.Close()

	byteValue, _ := ioutil.ReadAll(annotationFile)

	var annotationStruct Annotation
	err = json.Unmarshal(byteValue, &annotationStruct)
	if err != nil {
		return nil, err
	}

	result := []extractImagesInput{}

	catDict := make(map[int]int)
	for _, imageCat := range annotationStruct.Categories {
		num, err := strconv.Atoi(imageCat.Name)
		if err == nil {
			if num < 10 {
				catDict[imageCat.Id] = num
			}
		}
	}

	annDict := make(map[int][]coords)
	for _, imageAnn := range annotationStruct.Annotations {
		if num, exists := catDict[imageAnn.CategoryId]; exists {
			annDict[imageAnn.ImageId] = append(annDict[imageAnn.ImageId], coords{
				x:      int(imageAnn.Bbox[0]),
				y:      int(imageAnn.Bbox[1]),
				width:  int(imageAnn.Bbox[2]),
				height: int(imageAnn.Bbox[3]),
				number: num,
			})
		}
	}

	for _, image := range annotationStruct.Images {
		if _, exists := annDict[image.Id]; exists {
			result = append(result, extractImagesInput{
				imageName: image.FileName,
				coords:    annDict[image.Id],
			})
		}
	}

	return result, nil
}

type Annotation struct {
	Categories []struct {
		Id   int    `json:"id"`
		Name string `json:"name"`
	} `json:"categories"`
	Images []struct {
		Id       int    `json:"id"`
		FileName string `json:"file_name"`
	} `json:"images"`
	Annotations []struct {
		ImageId    int       `json:"image_id"`
		CategoryId int       `json:"category_id"`
		Bbox       []float64 `json:"bbox"`
	} `json:"annotations"`
}
