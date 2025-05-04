package core

import (
	"fmt"
	"gopkg.in/yaml.v3"
	"os"
	"rbac_manager/config"
)

func GetConfig() *config.Config {
	bytes, err := os.ReadFile("config/config.yaml")
	if err != nil {
		fmt.Printf("read config file err: %v", err)
		return nil
	}
	cfg := new(config.Config)
	err = yaml.Unmarshal(bytes, cfg)
	if err != nil {
		fmt.Printf("config file form err: %v", err)
		return nil
	}
	return cfg
}
