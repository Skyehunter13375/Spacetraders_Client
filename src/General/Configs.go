package General

import "os"
import "gopkg.in/yaml.v3"

type Configs struct {
	DB struct {
		DbPath  string `yaml:"DbPath"`
		DbBuild string `yaml:"DbBuild"`
		DbReset string `yaml:"DbReset"`
	} `yaml:"DB"`

	LOG struct {
		ErrPath string `yaml:"ErrPath"`
		ActPath string `yaml:"ActPath"`
	} `yaml:"LOG"`

	API struct {
		AgentName  string `yaml:"AgentName"`
		AccntToken string `yaml:"AccntToken"`
		AgentToken string `yaml:"AgentToken"`
	} `yaml:"API"`
}

func GetConfig() (*Configs, error) {
	data, err := os.ReadFile("config.yaml")
	if err != nil { LogErr("GetConfig: " + err.Error()); return nil, err }

	var cfg Configs
	err = yaml.Unmarshal(data, &cfg)
	if err != nil { LogErr("GetConfig: " + err.Error()); return nil, err }

	return &cfg, nil
}
