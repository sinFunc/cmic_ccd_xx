package config

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
)

type AppConfig struct {
	Logger struct {
		File  string
		Level string
	}
	Server struct {
		Ip   string
		Port int
	}
}

func (c *AppConfig) Init(file string) (err error) {
	var data []byte
	data, err = ioutil.ReadFile(file)
	if err != nil {
		return
	}

	if err = yaml.Unmarshal(data, c); err != nil {
		return
	}

	err = initLogger()

	return
}

//singleton global appConfig handler
//var pAppConfig *AppConfig
//
//func GetAppConfig() *AppConfig {
//	return pAppConfig
//}

//func InitAppConfig(file string) (err error) {
//	if pAppConfig != nil {
//		return errors.New("already")
//	}
//	pAppConfig = new(AppConfig)
//
//	var data []byte
//	data, err = ioutil.ReadFile(file)
//	if err != nil {
//		return
//	}
//
//	if err = yaml.Unmarshal(data, pAppConfig); err != nil {
//		return
//	}
//
//	err = initLogger()
//
//	return
//
//}
