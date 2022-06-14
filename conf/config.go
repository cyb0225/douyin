// 加载、读取用户配置的config文件


package config

import (
	"io/ioutil"
	"log"

	"gopkg.in/yaml.v2"
)

type DataBaseConfig struct {
	Username string `yaml:"username"`
	Password string `yaml:"password"`
	Host     string `yaml:"host"`
	Port     int    `yaml:"port"`
	Database string `yaml:"database"`
}

type ObjectStorageConfig struct {
	Endpoint        string `yaml:"endpoint"`
	AccessKeyId     string `yaml:"accessKeyId"`
	AccessKeySecret string `yaml:"accessKeySecret"`
	Examplebucket   string `yaml:"examplebucket"`
}

var (
	DBconf DataBaseConfig
	OSconf ObjectStorageConfig
)

func InitConfig() error {
	configFile, err := ioutil.ReadFile("./conf/config.yaml")
	if err != nil {
		return err
	}
	// 读取mysql数据库信息
	if err := yaml.Unmarshal(configFile, &DBconf); err != nil {
		return err
	}

	// 读取oss对象存储信息
	if err := yaml.Unmarshal(configFile, &OSconf); err != nil {
		return err
	}

	log.Println("configured successfully")
	return nil
}
