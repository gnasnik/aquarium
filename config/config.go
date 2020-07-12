package config

import (
	"encoding/json"
	"io/ioutil"
	"os"

	"github.com/pkg/errors"
)

var Configs *Config

type Config struct {
	RunMode       string `json:"run_mode"`
	Host          string `json:"host"`
	Port          int    `json:"port"`
	JwtUserSecret string `json:"jwt_user_secret"`
	DBURL         string `json:"dburl"`
}

func Exist(filename string) bool {
	_, err := os.Stat(filename)
	return err == nil || os.IsExist(err)
}

func InitConfig(file string) error {
	if !Exist(file) {
		return errors.New("no config file")
	}
	in, err := ioutil.ReadFile(file)
	if err != nil {
		return err
	}

	Configs = new(Config)
	if err := json.Unmarshal(in, Configs); err != nil {
		return err
	}

	if err := initDB(); err != nil {
		return err
	}

	return nil
}
