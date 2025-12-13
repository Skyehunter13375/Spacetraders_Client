package Task

import "os"
import "gopkg.in/yaml.v3"
import "Spacetraders/src/Model"

// FEAT: Get existing config YAML file
func GetConfig() (*Model.Configs, error) {
	data, err := os.ReadFile("config.yaml")
	if err != nil { LogErr("GetConfig: " + err.Error()); return nil, err }

	var cfg Model.Configs
	err = yaml.Unmarshal(data, &cfg)
	if err != nil { LogErr("GetConfig: " + err.Error()); return nil, err }

	return &cfg, nil
}

// FEAT: Store updated config YAML file
func SaveConfig(cfg *Model.Configs) error {
	data, err := yaml.Marshal(cfg)
	if err != nil { LogErr("SaveConfig: Marshal failed: " + err.Error()) }

	err = os.WriteFile("config.yaml", data, 0644)
	if err != nil { LogErr("SaveConfig: Write failed: " + err.Error()) }

	return nil
}

