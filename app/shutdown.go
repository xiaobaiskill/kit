package app

import (
	"fmt"
	"gitee.com/jinmingzhi/kit/eventbus"
	"log"
	"os"
	"os/signal"
	"syscall"
)

const (
	ExitTopic = "app.exit"
)

var (
	signCh = make(chan os.Signal)
	exitCh = make(chan struct{})
)

func init() {
	eventbus.SubscribeFunc(ExitTopic, func(msg interface{}) {
		log.Println(msg)
		exitCh <- struct{}{}
	})
	signal.Notify(signCh, syscall.SIGINT, syscall.SIGTERM)
}

func Exit() {
	sig := <-signCh
	Existing(fmt.Sprintf("receive signal %v", sig))
	os.Exit(0)
}

func Existing(msg interface{}) {
	eventbus.Publish(ExitTopic, msg)
	<-exitCh
}
