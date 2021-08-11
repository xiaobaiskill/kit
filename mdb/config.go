package mdb

import "gitee.com/jinmingzhi/kit/util"

type Config struct {
	URL string
}

func MustLoadConfig() *Config {
	cfg := new(Config)
	util.MustLoadConfig(cfg)
	return cfg
}
