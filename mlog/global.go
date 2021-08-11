package mlog

import (
	"github.com/xiaobaiskill/kit/app"
	"github.com/xiaobaiskill/kit/eventbus"
	"io"
	"log"
	"sync"
)

var (
	std       = New(MustLoadConfig())
	loggerMap sync.Map
)

func init() {
	eventbus.SubscribeFunc(app.ExitTopic, func(msg interface{}) {
		Sync()
		log.Println("all loggers synced")
	})
}

func Logger(name string) *MLog {
	if name == "" {
		return std
	}

	logger, ok := loggerMap.Load(name)
	if !ok {
		logger = std.Clone(name)
		loggerMap.Store(name, logger)
	}

	return logger.(*MLog)
}

func Sync() {
	std.Sync()
	loggerMap.Range(func(key, value interface{}) bool {
		l := value.(*MLog)
		l.Sync()
		return true
	})
}

func Ignore(err error) {
	std.Ignore(err)
}

func Close(c io.Closer) {
	std.Close(c)
}
