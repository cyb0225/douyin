// // 加载使用config文件
package config

import (
	"io/ioutil"
	"log"

	"gopkg.in/yaml.v2"
)

type DataBaseConfig struct {
	Username string `yaml:"username"`
	Password string `yaml:"password"`
	Host       string `yaml:"host"`
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

func InitConfig() error{
	configFile, err := ioutil.ReadFile("./config/config.yaml")
	if err != nil {
		return err
	}
	if err := yaml.Unmarshal(configFile, &DBconf); err != nil {
		return err
	}

	if err := yaml.Unmarshal(configFile, &OSconf); err != nil {
		return err
	}
	log.Println(DBconf)
	log.Println(OSconf)
	log.Println("configured successfully")
	return nil
}
