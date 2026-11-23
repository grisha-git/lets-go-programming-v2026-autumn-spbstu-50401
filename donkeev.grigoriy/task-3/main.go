package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"os"
	"path/filepath"

	"github.com/donkeev.grigoriy/task-3/config"
	"github.com/donkeev.grigoriy/task-3/parser"
)

func main() {
	configPath := flag.String("config", "config.yaml", "Path to config file")
	flag.Parse()
	if flag.NArg() > 0 {
		panic(fmt.Sprintf("Too many arguments"))
	}

	configVal, err := config.Load(*configPath)
	if err != nil {
		panic(fmt.Sprintf("Download error"))
	}

	xmlData, err := os.ReadFile(configVal.InputFile)
	if err != nil {
		panic(fmt.Sprintf("Reading error"))
	}
	currencyData, err := parser.XMLFormatParser(xmlData)
	if err != nil {
		panic(fmt.Sprintf("Parsing to XML error: %v", err))
	}

	outputDir := filepath.Dir(configVal.OutputFile)

	if err := os.MkdirAll(outputDir, 0755); err != nil {
		panic(fmt.Sprintf("Creating directory error"))
	}

	jsonData, err := json.MarshalIndent(currencyData, "", "    ")
	if err != nil {
		panic(fmt.Sprintf("Coding to json error %v", err))
	}

	err = os.WriteFile(configVal.OutputFile, jsonData, 0644)
	if err != nil {
		panic(fmt.Sprintf("Error file writing"))
	}
}
