package parser

import (
	"encoding/xml"
	"sort"
	"strconv"
	"strings"

	"github.com/donkeev.grigoriy/task-3/models"
)

func XMLFormatParser(data []byte) ([]models.OutputCurrency, error) {
	var valCurs models.ValCurs

	err := xml.Unmarshal(data, &valCurs)
	if err != nil {
		return nil, err
	}

	result := make([]models.OutputCurrency, 0, len(valCurs.Valutes))

	for _, v := range valCurs.Valutes {
		numCode, err := strconv.Atoi(v.NumCode)
		if err != nil {
			return nil, err
		}

		valueStr := strings.Replace(v.Value, ",", ".", -1)
		value, err := strconv.ParseFloat(valueStr, 64)
		if err != nil {
			return nil, err
		}

		result = append(result, models.OutputCurrency{
			NumCode:  numCode,
			CharCode: v.CharCode,
			Value:    value,
		})
	}

	sort.Slice(result, func(i, j int) bool {
		return result[i].Value > result[j].Value
	})

	return result, nil
}
