package config

import (
	"github.com/BurntSushi/toml"
	"path"
	"path/filepath"
	"runtime"
)

func InitConfig(reLayerAddress string) {
	currentAbPath := getCurrentAbsPathByCaller()
	tomlFile, err := filepath.Abs(currentAbPath + "/config.toml")
	if err != nil {
		panic("read toml file err: " + err.Error())
		return
	}

	if _, err := toml.DecodeFile(tomlFile, &Config); err != nil {
		panic("read toml file err: " + err.Error())
		return
	}

	ParseChainConfig(reLayerAddress)
}

func getCurrentAbsPathByCaller() string {
	var abPath string
	_, filename, _, ok := runtime.Caller(0)
	if ok {
		abPath = path.Dir(filename)
	}
	return abPath
}
