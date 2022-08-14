package config

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"sync"

	"github.com/VrMolodyakov/url-shortener/pkg/logging"
	"github.com/ilyakaznacheev/cleanenv"
)

type Config struct {
	LogLvl string `yaml : "loglvl"`
	Redis  Redis  `yaml : "redis"`
}

type Redis struct {
	Host     string `json:"host"`
	Port     string `json:"port"`
	Password string `json:"password"`
	DbNumber int    `json:"dbnumber"`
}

var instance *Config
var once sync.Once

func GetConfig() *Config {
	once.Do(func() {
		path, _ := os.Getwd()
		fmt.Println("path:", path)
		root := filepath.Dir(filepath.Dir(path))
		fmt.Println("dir2:", root)
		instance = &Config{}
		logger := logging.GetLogger("info")
		logger.Info("start config initialisation")
		configPath := root + "\\config\\config.yaml"
		dockerPath, _ := filepath.Abs(filepath.Dir(os.Args[0]))
		if exist, _ := Exists(configPath); exist {
			if err := cleanenv.ReadConfig(root+"\\config\\config.yaml", instance); err != nil {
				help, _ := cleanenv.GetDescription(instance, nil)
				logger.Info(help)
				logger.Fatal(err)
			}
		} else if exist, _ := Exists(dockerPath + "/config/config.yaml"); exist {
			if err := cleanenv.ReadConfig(dockerPath+"/config/config.yaml", instance); err != nil {
				help, _ := cleanenv.GetDescription(instance, nil)
				logger.Info(help)
				logger.Fatal(err)
			}
		}

	})
	return instance
}

func Exists(name string) (bool, error) {
	_, err := os.Stat(name)
	if err == nil {
		return true, nil
	}
	if errors.Is(err, os.ErrNotExist) {
		return false, nil
	}
	return false, err
}
