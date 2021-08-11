### 简介
```
日志封装
```

### 使用
* 设置日志 MOd
```
// local or server
// local 本地使用     
// server 上线使用    默认不含有debug,且输出 不含追踪信息
export LOG_MODE="local"    
```

* 关于使用
```
vat log = mlog.Logger("module name")
log.Debug()
log.Info()
log.Warn()
log.Fatal() // panic

// 参数
zap.String()
zap.Any()
zap.Error()
...
```

* 加入路由
```
Route(mux)
```

* 通过http控制模块的调整某一个module_name的级别
```
curl -X GET http://<ip>/logs/<module_name>/level
curl -X PUT -d '{"level":"warn"}' http://<ip>/logs/<module_name>/level
```
