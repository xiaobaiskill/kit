### 简介
```
随机数生成器
```

### 使用
* RandStringBytesMaskImprSrcUnsafe
```
// 获取随机数 长度4
str := RandStringBytesMaskImprSrcUnsafe(4)
```

* recovery_code
```
// 获取recovery_code
strs := GenerateRecoveryCode(5, 2) []string
fmt.Println(strs)
// [aa27r-8k2z2 50gw2-ohyqz]
```

