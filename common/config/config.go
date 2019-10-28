package config

// Config is the exported global constant for server configurations
var Config = &struct {
	Server   ServerConfig `yaml:"server"`
	Database MysqlConfig  `yaml:"database"`
}{}

// ServerConfig ...
type ServerConfig struct {
	Port       int32
	ValidToken string `yaml:"valid_token"`
	MaxSizeUpload    int    `yaml:"max_size"`
}

// MysqlConfig ...
type MysqlConfig struct {
	Uri     string `yaml:"uri"`
	MaxIdle int    `yaml:"maxIdle"`
	MaxOpen int    `yaml:"maxOpen"`
	ShowSql bool   `yaml:"showSql"`
}