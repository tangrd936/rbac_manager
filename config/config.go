package config

type Config struct {
	Db     Db     `yaml:"db"`
	System System `yaml:"system"`
}
type Db struct {
	Host     string `yaml:"host"`
	Password string `yaml:"password"`
	Port     int    `yaml:"port"`
	Username string `yaml:"username"`
}
type System struct {
	IP   string `yaml:"ip"`
	Port int    `yaml:"port"`
}
