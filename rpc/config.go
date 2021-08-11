package rpc

type Config struct {
	ListenAddress string
}

var defaultCfg = &Config{
	ListenAddress: ":50051",
}
