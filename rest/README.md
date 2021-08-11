### 简介
```
网关统一启动程序

默认已做了报错统一处理 和 开启了允许跨域请求
```

### 使用
* 获取rpc实例
```
rest := NewDefaultServer()
or 
rest := NewServer()
```

* 启动rest server
```
rest.ListenAndServed
```


* 编写中间件
```
# options 是要来加中间件的。 请将中间件的写法写成
type MiddleWareHeadle func(handler http.Handler) http.Handler

rest := NewServerWithConfig(options ...OptionHeadle)


例如: 简单的测速 中间件
func Speed(h http.Handler) http.Handler {
    return  http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        start := time.Now()
        h.ServeHTTP(w, r)
        log.Info("speed:", time.Now().Sub(start).Milliseconds())    
    })
}
```

* 添加中间件
```
rest := NewDefaultServer(...MiddleWareHeadle)
```


### 特别说明
```
1、默认会有/health 接口，用于检测接口是否正常

```
