package General

import (
	"os"

	"gopkg.in/yaml.v3"
)

type Configs struct {
	DB struct {
		Type string `yaml:"type"`
		Name string `yaml:"name"`
		Host string `yaml:"host"`
		Port int32  `yaml:"port"`
		User string `yaml:"user"`
		Pass string `yaml:"pass"`
		SSL  string `yaml:"SSL"`
	} `yaml:"database"`
	Tokens struct {
		Account string `yaml:"accnt"`
		Agent   string `yaml:"agent"`
	} `yaml:"tokens"`
}

func GetConfig() (*Configs, error) {
	data, err := os.ReadFile("config.yaml")
	if err != nil { LogErr("GetConfig: " + err.Error()); return nil, err }

	var cfg Configs
	err = yaml.Unmarshal(data, &cfg)
	if err != nil { LogErr("GetConfig: " + err.Error()); return nil, err }

	return &cfg, nil
}
