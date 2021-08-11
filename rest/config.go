package rest

type Config struct {
	CORSMaxAge    int
	ListenAddress string
	EmitDefaults  bool
}

var defaultCfg = &Config{
	CORSMaxAge:    86400,
	ListenAddress: ":80",
	EmitDefaults:  true,
}
