### 简介
```
env
环境变量 获取至结构体中
```

### 使用
* 定义 struct
```
type Student struct {
	Name   string   `env:"STUDENT_NAME,required"`
	Class  string   `env:"STUDENT_CLASS" envDefault:"5-1"`
	Gender int      `env:"STUDENT_GENDET" envDefault:"1"`
	Hobby  []string `env:"STUDENT_HOBBY" envSeparator:":"` //  default envSeparator:","
}
```
* 使用
```
export STUDENT_NAME="jmz"
export STUDENT_HOBBY="Basketball:book:go fishing"
s := &Student{}
MustLoadConfig(s)
fmt.Println(s)
```
### refer