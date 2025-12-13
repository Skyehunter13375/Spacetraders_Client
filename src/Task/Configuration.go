package Task

import "os"
import "gopkg.in/yaml.v3"
import "Spacetraders/src/Model"

func GetConfig() (*Model.Configs, error) {
	data, err := os.ReadFile("config.yaml")
	if err != nil { LogErr("GetConfig: " + err.Error()); return nil, err }

	var cfg Model.Configs
	err = yaml.Unmarshal(data, &cfg)
	if err != nil { LogErr("GetConfig: " + err.Error()); return nil, err }

	return &cfg, nil
}
