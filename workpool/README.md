### 简介
```
workpool
限流器
```

### 使用
* 开启调度
```
# work_num 任务并发数（工作数） 
StartDispathcher(work_num int)
```

* 加入job
```
# job 接口
WorkQueue <- job
```

* 关闭调度
```
StopDispathcher()
```
