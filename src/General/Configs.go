package General

import "os"
import "gopkg.in/yaml.v3"


func GetConfig() (*Configs, error) {
	data, err := os.ReadFile("config.yaml")
	if err != nil { LogErr("GetConfig: " + err.Error()); return nil, err }

	var cfg Configs
	err = yaml.Unmarshal(data, &cfg)
	if err != nil { LogErr("GetConfig: " + err.Error()); return nil, err }

	return &cfg, nil
}
