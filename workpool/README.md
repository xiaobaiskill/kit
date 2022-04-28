### 简介

```
workpool
限流器
```

### 使用

- 开启调度

```
# work_num 任务并发数（工作数）
workQueue := StartDispathcher(work_num int)
```

- 加入 job

```
# job 接口
workQueue <- job
```

- 关闭调度

```
StopDispathcher()
```
