package config

import (
	"os"

	"gopkg.in/yaml.v3"
)

type Config struct {
	InputFile  string `yaml:"input-file"`
	OutputFile string `yaml:"output-file"`
}

func Load(path string) (*Config, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	var configVal Config
	if err := yaml.Unmarshal(data, &configVal); err != nil {
		return nil, err
	}

	return &configVal, nil
}
