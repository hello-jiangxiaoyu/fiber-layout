package conf

import (
	"os"

	"gopkg.in/yaml.v3"
)

type SystemConfig struct {
	Svc        Server   `json:"server" yaml:"server"`
	Log        Log      `json:"log" yaml:"log"`
	Clickhouse Database `json:"clickhouse" yaml:"clickhouse"`
	DB         Database `json:"database" yaml:"database"`
	Redis      Database `json:"redis" yaml:"redis"`
}

func NewSystemConfig() (*SystemConfig, error) {
	yamlFile, err := os.ReadFile(GetSystemConfigFileName())
	if err != nil {
		return nil, err
	}

	var res SystemConfig
	if err = yaml.Unmarshal(yamlFile, &res); err != nil {
		return nil, err
	}

	return &res, nil
}
