# MySQL & Redis 迁移到 GORM 指南

## 项目结构

```
go-db-example/
├── config/
│   └── config.go          # 配置管理
├── database/
│   ├── mysql.go           # GORM MySQL 连接
│   └── redis.go           # go-redis 连接
├── models/
│   ├── user.go            # 用户模型
│   └── article.go         # 文章模型
├── repository/
│   ├── user_repository.go     # 用户数据访问层
│   └── article_repository.go  # 文章数据访问层
├── go.mod
└── main.go
```

---

## 1. MySQL 连接方式对比

### ❌ 旧方式：database/sql

```go
import (
    "database/sql"
    _ "github.com/go-sql-driver/mysql"
)

// 连接
db, err := sql.Open("mysql", "user:password@tcp(localhost:3306)/dbname")

// 查询
rows, err := db.Query("SELECT id, username, email FROM users WHERE status = ?", 1)
defer rows.Close()

var users []User
for rows.Next() {
    var u User
    rows.Scan(&u.ID, &u.Username, &u.Email)
    users = append(users, u)
}

// 插入
result, err := db.Exec(
    "INSERT INTO users (username, email, password) VALUES (?, ?, ?)",
    "alice", "alice@example.com", "hash",
)
id, _ := result.LastInsertId()

// 更新
_, err = db.Exec("UPDATE users SET nickname = ? WHERE id = ?", "Alice", 1)

// 事务
tx, _ := db.Begin()
tx.Exec(...)
tx.Commit() // or tx.Rollback()
```

### ✅ 新方式：GORM

```go
import (
    "gorm.io/driver/mysql"
    "gorm.io/gorm"
)

// 连接
dsn := "user:password@tcp(localhost:3306)/dbname?charset=utf8mb4&parseTime=True&loc=Local"
db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

// 查询
var users []User
db.Where("status = ?", 1).Find(&users)

// 插入
user := User{Username: "alice", Email: "alice@example.com", Password: "hash"}
db.Create(&user)  // user.ID 自动填充

// 更新
db.Model(&user).Update("nickname", "Alice")
db.Model(&User{}).Where("id = ?", 1).Updates(map[string]interface{}{"nickname": "Alice"})

// 事务
db.Transaction(func(tx *gorm.DB) error {
    if err := tx.Create(&user).Error; err != nil {
        return err  // 自动回滚
    }
    return nil  // 自动提交
})
```

---

## 2. 常用 GORM 操作对比

### 2.1 查询

| 操作 | database/sql | GORM |
|------|-------------|------|
| 单条查询 | `QueryRow(...).Scan(...)` | `db.First(&user, id)` |
| 条件查询 | `Query("... WHERE x=?", v)` | `db.Where("x = ?", v).Find(&users)` |
| 全部查询 | `Query("SELECT * FROM ...")` | `db.Find(&users)` |
| 统计 | `QueryRow("SELECT COUNT(*)")` | `db.Model(&User{}).Count(&count)` |

```go
// GORM 查询示例

// 主键查询
db.First(&user, 1)                           // SELECT * FROM users WHERE id = 1 LIMIT 1
db.First(&user, "id = ?", 1)                 // 同上

// 条件查询
db.Where("status = ?", 1).Find(&users)       // SELECT * FROM users WHERE status = 1
db.Where("name LIKE ?", "%jin%").Find(&users)
db.Where("created_at > ?", time.Now().Add(-24*time.Hour)).Find(&users)

// 多条件
db.Where("status = ? AND age > ?", 1, 18).Find(&users)
db.Where(map[string]interface{}{"status": 1, "role": "admin"}).Find(&users)

// Or 条件
db.Where("name = ?", "jinzhu").Or("name = ?", "jinzhu2").Find(&users)

// Not 条件
db.Not("status", 0).Find(&users)

// In 查询
db.Where("id IN ?", []int{1, 2, 3}).Find(&users)

// Between
db.Where("age BETWEEN ? AND ?", 18, 30).Find(&users)

// Select 指定字段
db.Select("id", "username").Find(&users)

// Order
db.Order("created_at DESC").Find(&users)
db.Order("age DESC, name").Find(&users)

// Limit & Offset
db.Offset(10).Limit(20).Find(&users)

// Group & Having
db.Model(&User{}).Select("role, count(*) as total").Group("role").Having("total > ?", 5).Find(&results)

// Distinct
db.Distinct("name").Find(&users)

// Join
db.Joins("JOIN orders ON orders.user_id = users.id").Find(&users)

// 子查询
subQuery := db.Model(&Order{}).Select("user_id").Where("amount > ?", 100)
db.Where("id IN (?)", subQuery).Find(&users)

// 原生 SQL
db.Raw("SELECT * FROM users WHERE status = ?", 1).Scan(&users)
```

### 2.2 创建

```go
// 单条创建
user := User{Username: "alice", Email: "alice@example.com"}
db.Create(&user)  // user.ID 自动填充

// 指定字段创建
db.Select("Username", "Email").Create(&user)

// 忽略字段
db.Omit("CreatedAt").Create(&user)

// 批量创建
users := []User{{Name: "a"}, {Name: "b"}, {Name: "c"}}
db.Create(&users)

// 分批创建（每次 100 条）
db.CreateInBatches(users, 100)

// Upsert（存在则更新）
db.Clauses(clause.OnConflict{
    Columns:   []clause.Column{{Name: "email"}},
    DoUpdates: clause.AssignmentColumns([]string{"nickname", "updated_at"}),
}).Create(&user)
```

### 2.3 更新

```go
// 保存所有字段（包括零值）
db.Save(&user)

// 更新单个字段
db.Model(&user).Update("nickname", "Alice")

// 更新多个字段
db.Model(&user).Updates(User{Nickname: "Alice", Age: 25})  // 零值不更新
db.Model(&user).Updates(map[string]interface{}{"nickname": "Alice", "age": 0})  // 零值也更新

// 条件更新
db.Model(&User{}).Where("status = ?", 0).Update("status", 1)

// 表达式更新
db.Model(&article).UpdateColumn("view_count", gorm.Expr("view_count + ?", 1))

// 批量更新
db.Model(&User{}).Where("role = ?", "user").Updates(map[string]interface{}{"status": 1})
```

### 2.4 删除

```go
// 软删除（需要 DeletedAt 字段）
db.Delete(&user, 1)                    // UPDATE users SET deleted_at = NOW() WHERE id = 1
db.Where("status = ?", 0).Delete(&User{})

// 硬删除
db.Unscoped().Delete(&user, 1)         // DELETE FROM users WHERE id = 1

// 查询包含软删除记录
db.Unscoped().Where("id = ?", 1).Find(&user)
```

---

## 3. 关联关系

### 3.1 一对多 (Has Many)

```go
// 模型定义
type User struct {
    ID       uint
    Name     string
    Articles []Article  // Has Many
}

type Article struct {
    ID       uint
    Title    string
    UserID   uint       // 外键
}

// 预加载
db.Preload("Articles").Find(&users)

// 条件预加载
db.Preload("Articles", "status = ?", 1).Find(&users)

// 嵌套预加载
db.Preload("Articles.Tags").Find(&users)

// 创建关联
db.Model(&user).Association("Articles").Append(&Article{Title: "new"})

// 删除关联
db.Model(&user).Association("Articles").Delete(&article)

// 清空关联
db.Model(&user).Association("Articles").Clear()

// 统计关联
count := db.Model(&user).Association("Articles").Count()
```

### 3.2 多对多 (Many to Many)

```go
// 模型定义
type Article struct {
    ID   uint
    Tags []Tag `gorm:"many2many:article_tags"`
}

type Tag struct {
    ID       uint
    Name     string
    Articles []Article `gorm:"many2many:article_tags"`
}

// 预加载
db.Preload("Tags").Find(&articles)

// 添加标签
db.Model(&article).Association("Tags").Append(&Tag{Name: "Go"})

// 替换标签
db.Model(&article).Association("Tags").Replace(newTags)
```

---

## 4. 事务

```go
// 方式 1: Transaction 闭包（推荐）
err := db.Transaction(func(tx *gorm.DB) error {
    if err := tx.Create(&user).Error; err != nil {
        return err  // 返回错误自动回滚
    }
    if err := tx.Create(&article).Error; err != nil {
        return err
    }
    return nil  // 返回 nil 自动提交
})

// 方式 2: 手动事务
tx := db.Begin()
defer func() {
    if r := recover(); r != nil {
        tx.Rollback()
    }
}()

if err := tx.Create(&user).Error; err != nil {
    tx.Rollback()
    return err
}
if err := tx.Create(&article).Error; err != nil {
    tx.Rollback()
    return err
}

tx.Commit()

// 嵌套事务（SavePoint）
db.Transaction(func(tx *gorm.DB) error {
    tx.Create(&user1)
    
    tx.Transaction(func(tx2 *gorm.DB) error {
        tx2.Create(&user2)
        return errors.New("rollback user2 only")  // 只回滚 user2
    })
    
    return nil  // user1 仍然提交
})
```

---

## 5. Redis 操作封装

### 连接

```go
import "github.com/redis/go-redis/v9"

rdb := redis.NewClient(&redis.Options{
    Addr:     "localhost:6379",
    Password: "",
    DB:       0,
    PoolSize: 10,
})

// 测试连接
ctx := context.Background()
_, err := rdb.Ping(ctx).Result()
```

### 常用操作

```go
ctx := context.Background()

// String
rdb.Set(ctx, "key", "value", 10*time.Minute)
val, _ := rdb.Get(ctx, "key").Result()

// Hash
rdb.HSet(ctx, "user:1", "name", "Alice", "age", "25")
rdb.HGet(ctx, "user:1", "name")
rdb.HGetAll(ctx, "user:1")

// List
rdb.LPush(ctx, "queue", "item1", "item2")
rdb.RPop(ctx, "queue")
rdb.LRange(ctx, "queue", 0, -1)

// Set
rdb.SAdd(ctx, "tags", "go", "rust")
rdb.SMembers(ctx, "tags")
rdb.SIsMember(ctx, "tags", "go")

// Sorted Set
rdb.ZAdd(ctx, "leaderboard", redis.Z{Score: 100, Member: "user1"})
rdb.ZRevRange(ctx, "leaderboard", 0, 9)  // Top 10

// 过期时间
rdb.Expire(ctx, "key", 1*time.Hour)
rdb.TTL(ctx, "key")

// 删除
rdb.Del(ctx, "key1", "key2")

// 自增
rdb.Incr(ctx, "counter")
rdb.IncrBy(ctx, "counter", 5)

// 分布式锁
locked, _ := rdb.SetNX(ctx, "lock:resource", "holder", 10*time.Second).Result()
if locked {
    defer rdb.Del(ctx, "lock:resource")
    // do something
}

// Pipeline（批量操作）
pipe := rdb.Pipeline()
pipe.Set(ctx, "key1", "val1", 0)
pipe.Set(ctx, "key2", "val2", 0)
pipe.Exec(ctx)

// 事务
rdb.Watch(ctx, func(tx *redis.Tx) error {
    _, err := tx.TxPipelined(ctx, func(pipe redis.Pipeliner) error {
        pipe.Set(ctx, "key", "value", 0)
        return nil
    })
    return err
}, "key")
```

---

## 6. 缓存策略示例

### Cache-Aside Pattern

```go
func (r *UserRepository) GetByID(ctx context.Context, id uint) (*User, error) {
    cacheKey := fmt.Sprintf("user:%d", id)
    
    // 1. 查缓存
    if cached, err := rdb.Get(ctx, cacheKey).Result(); err == nil {
        var user User
        json.Unmarshal([]byte(cached), &user)
        return &user, nil
    }
    
    // 2. 查数据库
    var user User
    if err := db.First(&user, id).Error; err != nil {
        return nil, err
    }
    
    // 3. 写缓存
    data, _ := json.Marshal(user)
    rdb.Set(ctx, cacheKey, data, 30*time.Minute)
    
    return &user, nil
}

func (r *UserRepository) Update(ctx context.Context, user *User) error {
    if err := db.Save(user).Error; err != nil {
        return err
    }
    // 删除缓存
    rdb.Del(ctx, fmt.Sprintf("user:%d", user.ID))
    return nil
}
```

---

## 7. 运行项目

```bash
# 设置环境变量
export MYSQL_HOST=localhost
export MYSQL_PORT=3306
export MYSQL_USER=root
export MYSQL_PASSWORD=your_password
export MYSQL_DATABASE=test

export REDIS_HOST=localhost
export REDIS_PORT=6379

# 安装依赖
cd go-db-example
go mod tidy

# 运行
go run main.go
```

---

## 8. 最佳实践

1. **连接池配置**：根据并发量配置 `MaxIdleConns` 和 `MaxOpenConns`
2. **超时设置**：所有外部调用都应设置超时
3. **错误处理**：区分 `gorm.ErrRecordNotFound` 和其他错误
4. **软删除**：使用 `gorm.DeletedAt` 字段实现
5. **缓存失效**：更新/删除时务必清除相关缓存
6. **批量操作**：大量数据使用 `CreateInBatches`
7. **预加载**：避免 N+1 问题，使用 `Preload`
8. **事务**：使用 `Transaction` 闭包自动管理
