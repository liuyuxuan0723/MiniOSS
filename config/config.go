package config

import (
	"fmt"
	"gopkg.in/yaml.v3"
	"os"
)

type Config struct {
	MySQl struct {
		Username string `yaml:"username"`
		Password string `yaml:"password"`
		Host     string `yaml:"host"`
		Port     string `yaml:"port"`
		Database string `yaml:"db"`
	} `yaml:"MySQL"`
	Server string `yaml:"Server"`
	Port   int    `yaml:"Port"`
}

var MyConfig Config

func LoadConfig() error {
	// 打开YAML文件
	path := "/Users/liuyuxuan/Documents/workbench/awesomeProject/config/config.yaml"
	file, err := os.Open(path)
	if err != nil {
		fmt.Printf("Failed to open file: %v\n", err)
		return err
	}
	defer file.Close()

	// 使用yaml解码器解析YAML文件，并将解析后的数据存储在config对象中
	decoder := yaml.NewDecoder(file)
	if err := decoder.Decode(&MyConfig); err != nil {
		fmt.Printf("Failed to decode YAML: %v\n", err)
		return err
	}
	return nil
}
