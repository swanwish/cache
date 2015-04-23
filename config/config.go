package config

import (
	"errors"

	"github.com/swanwish/go-helper/utils"
)

type Configuration interface {
	Load(filePath string) error
	Get(key string) string
}

var (
	Config Configuration = &IniConfiguration{}
)

var (
	ErrConfigurationTypeNotSpecified = errors.New("Configuration type not specified")
	ErrConfigurationFileNotExists    = errors.New("Configuration file not exists")
)

func Load(filePath string) error {
	if Config == nil {
		return ErrConfigurationTypeNotSpecified
	}
	if !utils.FileExists(filePath) {
		return ErrConfigurationFileNotExists
	}
	return Config.Load(filePath)
}

func Get(key string) (string, error) {
	if Config == nil {
		return "", ErrConfigurationTypeNotSpecified
	}
	return Config.Get(key), nil
}

//func NewConfig(configType, filePath string) (Configuration, error) {
//	var c Configuration
//	switch configType {
//	case "ini":
//		c = &IniConfiguration{}
//	}
//	if c != nil {
//		err := c.Load(filePath)
//		if err != nil {
//			return c, err
//		}
//		return c, nil
//	}
//	return nil, ErrConfigurationFileNotExists
//}
