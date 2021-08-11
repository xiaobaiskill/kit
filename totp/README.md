### 简介
```
totp
google 验证 生成器
```

### 包使用方式

#### 生成google 验证器
* Generate (生成google 验证器)
```
key,err := NewKey()
```

* key 生成 google验证器url
```
url := key.URL()
```


### 通过url验证用户code
* url 转成 key
```
key,err := NewKeyFromUrl(url)
```

* 验证 google 验证器内的 code
```
totp.Validate(code, key.Secret())
```


### refer
```
	github.com/pquerna/otp
```