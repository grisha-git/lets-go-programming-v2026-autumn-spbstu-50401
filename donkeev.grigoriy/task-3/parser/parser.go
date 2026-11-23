package parser

import (
	"bytes"
	"encoding/xml"
	"io"
	"sort"
	"strconv"
	"strings"

	"golang.org/x/text/encoding/charmap"

	"github.com/donkeev.grigoriy/task-3/models"
)

func XMLFormatParser(data []byte) ([]models.OutputCurrency, error) {
	var valCurs models.ValCurs

	decoder := xml.NewDecoder(bytes.NewReader(data))
	decoder.CharsetReader = func(charset string, input io.Reader) (io.Reader, error) {
		if charset == "windows-1251" {
			return charmap.Windows1251.NewDecoder().Reader(input), nil
		}
		return input, nil
	}

	if err := decoder.Decode(&valCurs); err != nil {
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
