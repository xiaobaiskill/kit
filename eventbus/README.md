事件总线(发布订阅模式)
---

### 简介
```
事件总线是发布/订阅模式的实现，其中发布者发布数据，并且订阅者可以监听这些数据并基于这些数据作出处理。
这使发布者与订阅者松耦合。发布者将数据事件发布到事件总线，总线负责将它们发送给订阅者。
```

### 使用
* 实现订阅者接口
```
type Sub interface {
	Receive(interface{})
}
```

* 订阅
```
Subscribe(topic, sub)
SubscribeFunc(topic, f func(msg interface{}){
    // to do
})
```

* 发布
```
Publish(topic string, msg interface{})
```


### 解说
#### 订阅者(sub)
* 当发布时， 会向对应的topic 发送msg, receive 接收msg
```
type Sub interface {
	Receive(interface{})
}
```
#### 订阅管理者(node)
```
管理一个种类的 订阅者们， 
添加订阅者
发布
```
#### 总线(bus)
```
管理订阅管理者
```

### refer
* [Golang学习篇—实现简单的事件总线(发布订阅模式)](https://blog.csdn.net/finghting321/article/details/103394274)
* [写一个简单的eventBus（Pub/Sub）](https://www.cnblogs.com/Jun10ng/p/13173368.html)