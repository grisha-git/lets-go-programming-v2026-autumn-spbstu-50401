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

	cfg, err := config.Load(*configPath)
	if err != nil {
		panic(fmt.Sprintf("Download error %v", err))
	}

	xmlData, err := os.ReadFile(cfg.InputFile)
	if err != nil {
		panic(fmt.Sprintf("Reading error '%s': %v", cfg.InputFile, err))
	}

	currencies, err := parser.XMLFormatParser(xmlData)
	if err != nil {
		panic(fmt.Sprintf("Parsing to XML error: %v", err))
	}

	outputDir := filepath.Dir(cfg.OutputFile)

	if err := os.MkdirAll(outputDir, 0755); err != nil {
		panic(fmt.Sprintf("Creating directory error '%s': %v", outputDir, err))
	}

	jsonData, err := json.MarshalIndent(currencies, "", "    ")
	if err != nil {
		panic(fmt.Sprintf("Coding to json error %v", err))
	}

	err = os.WriteFile(cfg.OutputFile, jsonData, 0644)
	if err != nil {
		panic(fmt.Sprintf("Error file writing '%s': %v", cfg.OutputFile, err))
	}
}
