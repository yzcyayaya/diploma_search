package config

import (
	"embed"

	"gopkg.in/yaml.v3"
)

// 全局配置
var C Config

//go:embed config.yaml
var f embed.FS

type Config struct {
	Server      Server
	Meilisearch Meilisearch
}

type Server struct {
	Port string `yaml:"port"`
	Name string `yaml:"name"`
	Host string `yaml:"host"`
}

type Meilisearch struct {
	Host   string `yaml:"host"`
	Port   string `yaml:"port"`
	Apikey string `yaml:"apikey"`
}

func init() {
	serverInit()
}
func serverInit() {
	server, err := f.ReadFile("config.yaml")
	if err != nil {
		panic(any("读取配置文件出现错误,可能不存在！"))
	}
	err = yaml.Unmarshal(server, &C)
	if err != nil {
		panic(any("读取配置文件出现错误,转换失败！"))
	}
}
