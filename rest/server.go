package rest

import (
	"fmt"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"github.com/xiaobaiskill/kit/app"
	"github.com/xiaobaiskill/kit/rest/middleware"
	"go.uber.org/zap"
	"net/http"
)

type Server struct {
	ServeMux     *runtime.ServeMux
	Handler      http.Handler
	HttpServeMux *http.ServeMux
	Address      string
}

func NewDefaultServer(options ...middleware.MiddleWareHeadle) *Server {
	return NewServer(defaultCfg, options...)
}

func NewServer(cfg *Config, options ...middleware.MiddleWareHeadle) *Server {
	restMux := runtime.NewServeMux(
		runtime.WithMarshalerOption(runtime.MIMEWildcard, &runtime.JSONPb{OrigName: true, EmitDefaults: cfg.EmitDefaults}),
		runtime.WithProtoErrorHandler(CustomRESTErrorHandler),
		runtime.WithIncomingHeaderMatcher(CustomMatcher), // HTTP请求头到gRPC客户端元数据的映射

	)

	httpServeMux := http.NewServeMux()
	httpServeMux.Handle("/", restMux)
	health(httpServeMux)

	var handler http.Handler = httpServeMux

	for _, handle := range options {
		handler = handle(handler)
	}

	return &Server{
		ServeMux:     restMux,
		Handler:      handler,
		HttpServeMux: httpServeMux,
		Address:      cfg.ListenAddress,
	}
}

// To keep the the default mapping rule alongside
func CustomMatcher(key string) (string, bool) {
	return runtime.DefaultHeaderMatcher(key)
}

func (s *Server) ListenAndServed() {
	go func() {
		log.Info("start serving rest service", zap.String("address", s.Address))
		if err := http.ListenAndServe(s.Address, s.Handler); err != nil {
			app.Existing(fmt.Sprintf("rest start failed:%v", err))
		}
	}()
}
