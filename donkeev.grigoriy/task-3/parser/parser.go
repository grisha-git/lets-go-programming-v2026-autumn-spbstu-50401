package parser

import (
	"bytes"
	"encoding/xml"
	"fmt"
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
		return nil, fmt.Errorf("xml decode: %w", err)
	}

	result := make([]models.OutputCurrency, 0, len(valCurs.Valutes))

	for _, value := range valCurs.Valutes {
		numCode, err := strconv.Atoi(value.NumCode)
		if err != nil {
			return nil, fmt.Errorf("types conversion: %w", err)
		}

		valueStr := strings.ReplaceAll(value.Value, ",", ".")

		val, err := strconv.ParseFloat(valueStr, 64)
		if err != nil {
			return nil, fmt.Errorf("types conversion: %w", err)
		}

		result = append(result, models.OutputCurrency{
			NumCode:  numCode,
			CharCode: value.CharCode,
			Value:    val,
		})
	}

	sort.Slice(result, func(i, j int) bool {
		return result[i].Value > result[j].Value
	})

	return result, nil
}
