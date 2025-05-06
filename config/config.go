package config

type Config struct {
	Db     Db     `yaml:"db"`
	Redis  Redis  `yaml:"redis"`
	System System `yaml:"system"`
}
type Db struct {
	Host     string `yaml:"host"`
	Port     int    `yaml:"port"`
	Username string `yaml:"username"`
	Password string `yaml:"password"`
}
type Redis struct {
	Host     string `yaml:"host"`
	Port     int    `yaml:"port"`
	Password string `yaml:"password"`
}
type System struct {
	IP   string `yaml:"ip"`
	Port int    `yaml:"port"`
}
