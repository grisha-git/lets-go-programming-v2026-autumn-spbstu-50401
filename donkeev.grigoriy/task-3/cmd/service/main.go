package main

import (
	"encoding/json"
	"flag"
	"os"
	"path/filepath"

	"github.com/donkeev.grigoriy/task-3/config"
	"github.com/donkeev.grigoriy/task-3/parser"
)

func main() {
	configPath := flag.String("config", "config.yaml", "Path to config file")
	flag.Parse()

	if flag.NArg() > 0 {
		panic("Too many arguments")
	}

	configVal, err := config.Load(*configPath)
	if err != nil {
		panic("Download error")
	}

	xmlData, err := os.ReadFile(configVal.InputFile)
	if err != nil {
		panic("Reading error")
	}

	currencyData, err := parser.XMLFormatParser(xmlData)
	if err != nil {
		panic("Parsing to XML error: %v")
	}

	outputDir := filepath.Dir(configVal.OutputFile)
	if err := os.MkdirAll(outputDir, 0o755); err != nil {
		panic("Creating directory error")
	}

	jsonData, err := json.MarshalIndent(currencyData, "", "    ")
	if err != nil {
		panic("Coding to json error")
	}

	err = os.WriteFile(configVal.OutputFile, jsonData, 0o600)
	if err != nil {
		panic("Error file writing")
	}
}
