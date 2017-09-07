package main

import (
	"github.com/BurntSushi/toml"
)

type Config struct {
	WindowWidth  int
	WindowHeight int
	IsResizable  bool
	IsFullscreen bool
	WindowName   string
}

func ReadConfig() Config {
	var conf Config
	if _, err := toml.DecodeFile("./config.toml", &conf); err != nil {
		panic(err)
	}
	return conf
}
