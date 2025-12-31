package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"go-db-example/config"
	"go-db-example/database"
	"go-db-example/models"
	"go-db-example/repository"
)

func main() {
	// 加载配置
	cfg := config.Load()

	// 初始化数据库连接
	if err := database.InitMySQL(cfg.MySQL); err != nil {
		log.Fatalf("MySQL init failed: %v", err)
	}
	defer database.CloseMySQL()

	if err := database.InitRedis(cfg.Redis); err != nil {
		log.Fatalf("Redis init failed: %v", err)
	}
	defer database.CloseRedis()

	// 自动迁移
	if err := database.AutoMigrate(&models.User{}, &models.Article{}, &models.Tag{}); err != nil {
		log.Fatalf("Migration failed: %v", err)
	}
	log.Println("Database migration completed")

	// 示例操作
	ctx := context.Background()
	demoUserOperations(ctx)
	demoArticleOperations(ctx)
	demoRedisOperations(ctx)
}

func demoUserOperations(ctx context.Context) {
	fmt.Println("\n=== User Operations ===")
	userRepo := repository.NewUserRepository()

	// 创建用户
	user := &models.User{
		Username: "alice",
		Email:    "alice@example.com",
		Password: "hashed_password",
		Nickname: "Alice",
		Status:   1,
	}

	if err := userRepo.Create(user); err != nil {
		log.Printf("Create user failed: %v", err)
	} else {
		fmt.Printf("Created user: %+v\n", user)
	}

	// 查询（带缓存）
	found, err := userRepo.GetByID(ctx, user.ID)
	if err != nil {
		log.Printf("Get user failed: %v", err)
	} else {
		fmt.Printf("Found user: %+v\n", found)
	}

	// 再次查询（应该命中缓存）
	found2, _ := userRepo.GetByID(ctx, user.ID)
	fmt.Printf("Found user (from cache): %+v\n", found2)

	// 更新
	user.Nickname = "Alice Updated"
	if err := userRepo.Update(ctx, user); err != nil {
		log.Printf("Update user failed: %v", err)
	}

	// 分页查询
	users, total, _ := userRepo.List(1, 10, map[string]interface{}{
		"status": 1,
	})
	fmt.Printf("User list: total=%d, users=%+v\n", total, users)
}

func demoArticleOperations(ctx context.Context) {
	fmt.Println("\n=== Article Operations ===")
	articleRepo := repository.NewArticleRepository()

	// 创建标签
	tag := models.Tag{Name: "Go"}
	database.DB.FirstOrCreate(&tag, models.Tag{Name: "Go"})

	// 创建文章
	article := &models.Article{
		Title:    "Getting Started with GORM",
		Content:  "GORM is a fantastic ORM library for Go...",
		Summary:  "Learn GORM basics",
		AuthorID: 1,
		Status:   1,
		Tags:     []models.Tag{tag},
	}

	if err := articleRepo.Create(article); err != nil {
		log.Printf("Create article failed: %v", err)
	} else {
		fmt.Printf("Created article: %+v\n", article)
	}

	// 查询（带预加载）
	found, err := articleRepo.GetByID(ctx, article.ID)
	if err != nil {
		log.Printf("Get article failed: %v", err)
	} else {
		fmt.Printf("Found article with author: %+v\n", found)
	}

	// 增加浏览量
	_ = articleRepo.IncrViewCount(ctx, article.ID)
	_ = articleRepo.IncrViewCount(ctx, article.ID)
	count, _ := articleRepo.GetViewCount(ctx, article.ID)
	fmt.Printf("View count: %d\n", count)

	// 分页查询
	articles, total, _ := articleRepo.List(1, 10, map[string]interface{}{
		"status": 1,
	})
	fmt.Printf("Article list: total=%d, count=%d\n", total, len(articles))
}

func demoRedisOperations(ctx context.Context) {
	fmt.Println("\n=== Redis Operations ===")

	// String
	_ = database.Set(ctx, "greeting", "Hello, GORM!", 5*time.Minute)
	val, _ := database.Get(ctx, "greeting")
	fmt.Printf("String: %s\n", val)

	// Hash
	_ = database.HSet(ctx, "user:1", "name", "Bob", "age", "25")
	name, _ := database.HGet(ctx, "user:1", "name")
	fmt.Printf("Hash name: %s\n", name)

	all, _ := database.HGetAll(ctx, "user:1")
	fmt.Printf("Hash all: %+v\n", all)

	// List
	_ = database.LPush(ctx, "queue", "task1", "task2", "task3")
	tasks, _ := database.LRange(ctx, "queue", 0, -1)
	fmt.Printf("List: %v\n", tasks)

	// Set
	_ = database.SAdd(ctx, "tags", "go", "rust", "python")
	members, _ := database.SMembers(ctx, "tags")
	fmt.Printf("Set members: %v\n", members)

	// 分布式锁
	locked, _ := database.SetNX(ctx, "lock:resource", "holder1", 10*time.Second)
	fmt.Printf("Lock acquired: %v\n", locked)
}
