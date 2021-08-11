package main

import (
	"gitee.com/jinmingzhi/kit/app"
	"gitee.com/jinmingzhi/kit/mlog"
	"net/http"
	"time"
)

func main() {
	log1 := mlog.Logger("dev")
	log2 := mlog.Logger("stage")
	log3 := mlog.Logger("pro")
	var n int
	go func() {
		for {
			n++
			switch n % 3 {
			case 0:
				log1.Info("aaaaaaa")
			case 1:
				log2.Debug("bbbbb")
			case 2:
				log3.Error("ccccc")
			}
			time.Sleep(time.Millisecond * 500)
		}
	}()

	mux := http.NewServeMux()
	mlog.Route(mux)
	go http.ListenAndServe(":8080", mux)
	app.Exit()
}
