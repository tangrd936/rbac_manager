package config

type Config struct {
	Db     Db     `yaml:"db"`
	Redis  Redis  `yaml:"redis"`
	System System `yaml:"system"`
	Jwt    Jwt    `yaml:"jwt"`
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

type Jwt struct {
	Secret string `yaml:"secret"` // 密钥
	Expire int    `yaml:"expire"` // 过期时间
	Issuer string `yaml:"issuer"` // 颁发人
}
