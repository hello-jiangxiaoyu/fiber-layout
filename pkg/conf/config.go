package conf

import (
	"os"

	"gopkg.in/yaml.v3"
)

type (
	Server struct {
		Listen         string `json:"listen" yaml:"listen"`
		IsUpload       bool   `json:"uploadSwitch" yaml:"uploadSwitch"`
		IsWildcard     bool   `json:"wildcardSwitch" yaml:"wildcardSwitch"`
		WildcardSuffix string `json:"wildcardSuffix" yaml:"wildcardSuffix"`
		SessionName    string `json:"sessionName" yaml:"sessionName"`
		StoreKeyPairs  string `json:"storeKeyPairs" yaml:"storeKeyPairs"`
	}

	Log struct {
		Level       string `json:"level" yaml:"level"`
		Dir         string `json:"dir" yaml:"dir"`
		IsFullStack bool   `json:"fullStackSwitch" yaml:"fullStackSwitch"`
	}

	SystemConfig struct {
		Svc Server   `json:"server" yaml:"server"`
		Log Log      `json:"log" yaml:"log"`
		DB  Database `json:"database" yaml:"database"`
	}
)

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
