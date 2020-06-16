package env

import (
	"log"
	"os"
	"path/filepath"
)

var (
	Dir     string
	LogDir  string
	LogPath string
)

func init() {
	file, _ := filepath.Abs(os.Args[0])
	dir := filepath.Dir(file)

	Dir = filepath.Dir(dir + "..")
	LogDir = Dir + "/log/"

	if !IsExist(LogDir) {
		if err := os.MkdirAll(LogDir, os.ModePerm); err != nil {
			log.Fatal(err)
		}
	}

	LogPath = LogDir + filepath.Base(os.Args[0]) + ".log"
	LogDir, _ = filepath.Abs(LogDir)
	LogPath, _ = filepath.Abs(LogPath)

	println(LogPath)
}

func IsExist(path string) bool {
	if _, err := os.Stat(path); err != nil {
		if os.IsExist(err) {
			return true
		}
	}
	return false
}
