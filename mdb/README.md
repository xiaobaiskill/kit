mongo 
---

### 使用
* 初始化
```
cli = NewClient(&Config{URL: "mongodb://localhost:27017"})
repo := NewRepository(cli,"<database>","<table>")
```

* 查（duo）
```
result := make([]*model.AppStatus, 0)
day := time.Now().Add(-time.Minute * 5)
repo.Find(ctx,
		bson.M{
			"status": bson.M{"$in": []model.Status{
				model.StatusShutdwon,
				model.StatusUnHealth,
			}},
			"is_personal": isPersonal,
			"createdat": bson.M{"$lt": day},
		},
		&result)
```
* 查（单）
```
var appStatus = new(model.AppStatus)
repo.FindOne(ctx, bson.M{"app_id": app_id}, appStatus)
```

* 增
```
repo.AddOne(ctx, data)
```

* 删
```
repo.DeleteMany(ctx, bson.M{"app_id": app_id})
repo.DeleteOne(ctx, bson.M{"app_id": app_id})
```
* 改
```
repo.UpdateOne(ctx, bson.M{"app_id": req.AppID}, data)
```

* 增改
```
# 有则改之，无则加之
repo.SaveOne(ctx, bson.M{"app_id": req.AppID}, data)
```
