package mlog

import (
	"io"
	"log"
	"sync"

	"github.com/xiaobaiskill/kit/app"
	"github.com/xiaobaiskill/kit/eventbus"
	"go.uber.org/zap/zapcore"
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

func SetLevel(level zapcore.Level) {
	std.SetLevel(level)
	loggerMap.Range(func(key interface{}, value interface{}) bool {
		log := value.(*MLog)
		log.SetLevel(level)
		return true
	})
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
