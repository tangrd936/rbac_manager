package flags

import (
	"flag"
	"os"
)

type option struct {
	Db   bool
	Env  string
	Menu string
	Type string
}

var FlagOptions option

func init() {
	flag.BoolVar(&FlagOptions.Db, "db", false, "迁移数据库")
	flag.StringVar(&FlagOptions.Env, "f", "", "运行环境")
	flag.StringVar(&FlagOptions.Menu, "m", "", "菜单")
	flag.StringVar(&FlagOptions.Type, "t", "", "类型")
	flag.Parse()
}

func Run() {
	if FlagOptions.Db {
		AutoMigrate()
		os.Exit(0)
	}
	switch FlagOptions.Menu {
	case "user":
		user := new(User)
		switch FlagOptions.Type {
		case "create":
			user.CreateAdminUser()
			os.Exit(0)
		}
	}
}
