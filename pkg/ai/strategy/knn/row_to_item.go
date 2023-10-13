package knn

import (
	"strconv"

	"github.com/jan-hajek/ai/pkg/ai/csvx"
	"github.com/pkg/errors"
)

type Item struct {
	number         int
	sourceFileName string
	fieldsAvgs     []float64
}

func rowsToItems(trainData []csvx.Row) (items []Item, _ error) {
	for _, row := range trainData {
		item, err := rowToItem(row)
		if err != nil {
			return nil, err
		}
		items = append(items, *item)
	}

	return
}

func rowToItem(row csvx.Row) (*Item, error) {
	number, err := strconv.Atoi(row[0])
	if err != nil {
		return nil, errors.WithStack(err)
	}

	avgs, err := convertStringsIntoFloats(row[2:])
	if err != nil {
		return nil, err
	}
	return &Item{
		number:         number,
		sourceFileName: row[1],
		fieldsAvgs:     avgs,
	}, nil
}

func convertStringsIntoFloats(input []string) ([]float64, error) {
	result := make([]float64, len(input))
	for i, value := range input {
		var err error
		result[i], err = strconv.ParseFloat(value, 64)
		if err != nil {
			return nil, errors.WithStack(err)
		}
	}
	return result, nil
}
