package core

import (
	"fmt"
	"gopkg.in/yaml.v3"
	"os"
	"rbac_manager/config"
	"rbac_manager/global"
)

func InitConfig(env string) {
	if env == "" {
		env = "dev"
	}
	cfgFile := "config/" + env + "_config.yaml"
	bytes, err := os.ReadFile(cfgFile)
	if err != nil {
		global.Log.Error(fmt.Sprintf("read config file err: %v", err))
		return
	}
	cfg := new(config.Config)
	err = yaml.Unmarshal(bytes, cfg)
	if err != nil {
		global.Log.Error(fmt.Sprintf("config file form err: %v", err))
		return
	}
	global.Log.Info(fmt.Sprintf("load config file success, config: %v", cfg))
	global.Conf = cfg
}
