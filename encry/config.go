package encry

import (
	"gitee.com/jinmingzhi/kit/util"
	"sync"
)


var (
	once sync.Once
	out *Config
)

type Config struct {
	SecretKey string `env:"ENCRY_SECRETKEY" envDefault:"MingzhiIsGoodMan"`
}

func MustLoadConfig() *Config {
	once.Do(func(){
		out = new(Config)
		util.MustLoadConfig(out)
	})

	return out
}
