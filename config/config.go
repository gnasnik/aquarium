package config

import (
	"encoding/json"
	"io/ioutil"
	"os"

	"github.com/pkg/errors"
)

var (
	RunMode       string
	JwtUserSecret string
	Host          string
	Port          int
)

type Config struct {
	RunMode       string `json:"run_mode"`
	Host          string `json:"host"`
	Port          int    `json:"port"`
	JwtUserSecret string `json:"jwt_user_secret"`
}

func Exist(filename string) bool {
	_, err := os.Stat(filename)
	return err == nil || os.IsExist(err)
}

func InitConfig() error {
	configJson := "config.json"
	if !Exist(configJson) {
		return errors.New("no config file")
	}
	in, err := ioutil.ReadFile(configJson)
	if err != nil {
		return err
	}
	var cfg = new(Config)
	if err := json.Unmarshal(in, cfg); err != nil {
		return err
	}
	RunMode = cfg.RunMode
	Host = cfg.Host
	Port = cfg.Port
	JwtUserSecret = cfg.JwtUserSecret
	return nil
}
