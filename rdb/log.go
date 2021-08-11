package rdb

import (
	"gitee.com/jinmingzhi/kit/mlog"
	"strings"
)

var (
	log = mlog.Logger("rdb")
)

func LogSQL(sql string, args ...interface{}) {
	outSQL := strings.ReplaceAll(sql, "?", "%v")
	log.Sugar().Debugf(outSQL, args...)
}
