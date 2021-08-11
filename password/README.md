### 简介
```
简单的passwd 效验器
```

### 使用
* 设置全局变量
```
# 加密key
export PASS_KEY="key"

# 加密后的token
export PASS_TOKEN="6b657976a2173be6393254e72ffa4d6df1030a"
```

* 验证
```
验证加密后的passwd 的token 是否和 token 一致，一致则正确。
Verfiy("passwd")
```

