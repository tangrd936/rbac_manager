package config

type Config struct {
	Db      Db      `yaml:"db"`
	Redis   Redis   `yaml:"redis"`
	System  System  `yaml:"system"`
	Jwt     Jwt     `yaml:"jwt"`
	Captcha Captcha `yaml:"captcha"`
	Email   Email   `yaml:"email"`
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
type Captcha struct {
	Enable bool `yaml:"enable"` // 是否启用验证码
}

type Email struct {
	Host   string `yaml:"host"`
	Port   int    `yaml:"port"`
	User   string `yaml:"user"`
	Pass   string `yaml:"pass"`
	Enable bool   `yaml:"enable"` // 是否启用邮箱
}
