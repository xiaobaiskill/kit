package password

import (
	"crypto/md5"
	"encoding/hex"
	"gitee.com/jinmingzhi/kit/util"
)

type passwdConfig struct {
	Key   string `env:"PASS_KEY" envDefault:"key"`
	Token string `env:"PASS_TOKEN" envDefault:"6b657976a2173be6393254e72ffa4d6df1030a"`
}

func MustLoadConfig() *passwdConfig {
	out := new(passwdConfig)
	util.MustLoadConfig(out)
	return out
}

func Verfiy(passwd string) bool {
	ctx := md5.New()
	ctx.Write([]byte(passwd))
	return hex.EncodeToString(ctx.Sum([]byte(Get().Key))) == Get().Token
}
