package config

import (
	"encoding/json"
	"io/ioutil"
)

type Config struct {
	Mode          string     `json:"mode"`
	Port          int        `json:"port"`
	DiagLogConfig *LogConfig `json:"diag"`
	StatLogConfig *LogConfig `json:"stat"`
}

type LogConfig struct {
	Level      string `json:"level"`
	Filename   string `json:"filename"`
	MaxSize    int    `json:"maxsize"`
	MaxAge     int    `json:"max_age"`
	MaxBackups int    `json:"max_backups"`
	Compress   bool   `json:"compress"`
}

var Conf = new(Config)

func LoadConfig(filePath string) error {
	b, err := ioutil.ReadFile(filePath)
	if err != nil {
		return err
	}
	return json.Unmarshal(b, Conf)
}
