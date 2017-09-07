package conf

import (
	"github.com/BurntSushi/toml"
)

type Type struct {
	WindowWidth  int
	WindowHeight int
	IsResizable  bool
	IsFullscreen bool
	WindowName   string
}

func Read() Type {
	var conf Type
	if _, err := toml.DecodeFile("./conf.toml", &conf); err != nil {
		panic(err)
	}
	return conf
}
