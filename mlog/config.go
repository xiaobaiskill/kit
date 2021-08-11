package mlog

import (
	"gitee.com/jinmingzhi/kit/util"
)

type Config struct {
	Mode Mode `env:"LOG_MODE" envDefault:"local"`
}

func MustLoadConfig() *Config {
	out := new(Config)
	util.MustLoadConfig(out)
	return out
}
