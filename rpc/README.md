### 简介
```
grpc 统一启动程序
```

### 使用

* 获取rest 实例
```
rpc := NewDefaultServer()
or
rpc := NewServer(config)
```

* 启动grpc
```
rpc.MustListenAndServe()
```

### 说明
```
默认添加了 对请求字段的验证

对于请求字段的验证需实现 功能
type validator interface {
	Validate() error
}

本人用的是 github.com/envoyproxy/protoc-gen-validate
```

### 案列
```
https://gitee.com/jinmingzhi/rest_test
```