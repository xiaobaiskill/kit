### 简介
```
sqlx 封装
```

### 使用

* 实例化
```
repo := NewMySQLRepository(Config{
        		DSN:             "root:root@tcp(127.0.0.1:3306)/test?parseTime=true",
        		ConnMaxLifetime: "30m",
        		MaxIdleConns:    "10",
        		SetMaxOpenConns: "20",
        	})
```

* 简单使用
```
id,err := repo.AddOne(ctx,`INSERT INTO test.food (name) VALUES (?)`, "apple")


type Res struct{
    ID string `db:"id"`
    Time Time `db:"time"`
}

data := &Res{}
err := repo.FindAll(ctx,data,`SELECT id, time FROM test.time WHERE id = ?`, "1")
```

* 关于事务的使用
```

// 在 ctx 中添加事务
tx,err := repo.NewWriteTx(context.Background())
if err !=nil{
    return err
}
defer tx.Rollback()

ctx := ContextWithTx(context.Background(), tx)


id,err := repo.AddOne(ctx,`INSERT INTO test.food (name) VALUES (?)`, "apple")

// 对sqlx 封装是有限的， 可以通过更多的扩展
repo.GetSQLOp(ctx)
```

* 关于数据类型的
```
type Pay struct {
	ID         uint32   `db:"id"`
	UID        uint32   `db:"uid"`
	PayNum     string   `db:"pay_num"` // 账单号
	SchoolTime string   `db:"school_time"`
	BatchID    uint32   `db:"batch_id"` // 报考批次id
	Amount     float32  `db:"amount"`
	Status     bool     `db:"status"` //是否已支付
	CreateTime Time `db:"create_time"`
}

CREATE TABLE IF NOT EXISTS `pay` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `uid` int(11) NOT NULL,
  `pay_num` varchar(255) NOT NULL,
  `school_time` varchar(255) NOT NULL,
  `batch_id` int(11)  NOT NULL,
  `amount` float NOT NULL,
  `status` tinyint(1) NOT NULL,
  `create_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin;
```

* 自定义数据类型
```
自定义的数据类型一定要实现下面前两个接口，后三个接口是用于输出的和 json 序列化使用的

// mysql 数据转自定义类型数据
type Scanner interface {
	// Scan assigns a value from a database driver.
	//
	// The src value will be of one of the following types:
	//
	//    int64
	//    float64
	//    bool
	//    []byte
	//    string
	//    time.Time
	//    nil - for NULL values
	//
	// An error should be returned if the value cannot be stored
	// without loss of information.
	//
	// Reference types such as []byte are only valid until the next call to Scan
	// and should not be retained. Their underlying memory is owned by the driver.
	// If retention is necessary, copy their values before the next call to Scan.
	Scan(src interface{}) error
}

// 自定义类型数据转 mysql 基础类型数据
type Valuer interface {
	// Value returns a driver Value.
	// Value must not panic.
	Value() (Value, error)
}

//  用于终端输出的
type Stringer interface {
    String() string
}

//  用于json输出的
type Marshaler interface {
	MarshalJSON() ([]byte, error)
}

// 用于json 序列号的
type Unmarshaler interface {
	UnmarshalJSON([]byte) error
}


// mysql 与 golang 基础数据的对应关系
varchar char text     =>    []byte
int bigint tinyit      =>   int64   bool
float DOUBLE DECIMAL  =>    float64
datetime              =>    time.time    (?parseTime=true,  否则是[]uint8)
```


### 案列
```
https://gitee.com/jinmingzhi/cdce-service
```