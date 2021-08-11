package util

import "github.com/caarlos0/env"

func MustLoadConfig(in interface{}) {
	if err := env.Parse(in); err != nil {
		panic(err)
	}
}
