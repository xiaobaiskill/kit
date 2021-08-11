token verfiy 
---

### 简介
```

token 验证器
```

### 使用
* 实例
```
t:=NewTokenVerfiy("key")
```
* 生成token
```
token := t.GenerateToken("secret")
```
* 验证token
```cassandraql
t.Verfiy(token,secret)
```

* 验证
```
验证加密后的secret的token 是否和 token 一致，一致则正确。
Verfiy("secret")
```

