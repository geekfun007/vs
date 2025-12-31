package database

import (
	"context"
	"fmt"
	"log"
	"time"

	"go-db-example/config"

	"github.com/redis/go-redis/v9"
)

var RDB *redis.Client

// InitRedis 初始化 Redis 连接
func InitRedis(cfg config.RedisConfig) error {
	RDB = redis.NewClient(&redis.Options{
		Addr:         fmt.Sprintf("%s:%d", cfg.Host, cfg.Port),
		Password:     cfg.Password,
		DB:           cfg.DB,
		PoolSize:     10,              // 连接池大小
		MinIdleConns: 5,               // 最小空闲连接
		DialTimeout:  5 * time.Second, // 连接超时
		ReadTimeout:  3 * time.Second, // 读超时
		WriteTimeout: 3 * time.Second, // 写超时
	})

	// 测试连接
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := RDB.Ping(ctx).Err(); err != nil {
		return fmt.Errorf("failed to connect redis: %w", err)
	}

	log.Println("Redis connected successfully")
	return nil
}

// CloseRedis 关闭 Redis 连接
func CloseRedis() error {
	if RDB != nil {
		return RDB.Close()
	}
	return nil
}

// ============ Redis 常用操作封装 ============

// Set 设置键值（带过期时间）
func Set(ctx context.Context, key string, value interface{}, expiration time.Duration) error {
	return RDB.Set(ctx, key, value, expiration).Err()
}

// Get 获取值
func Get(ctx context.Context, key string) (string, error) {
	return RDB.Get(ctx, key).Result()
}

// Del 删除键
func Del(ctx context.Context, keys ...string) error {
	return RDB.Del(ctx, keys...).Err()
}

// Exists 检查键是否存在
func Exists(ctx context.Context, key string) (bool, error) {
	n, err := RDB.Exists(ctx, key).Result()
	return n > 0, err
}

// Expire 设置过期时间
func Expire(ctx context.Context, key string, expiration time.Duration) error {
	return RDB.Expire(ctx, key, expiration).Err()
}

// HSet 设置 Hash 字段
func HSet(ctx context.Context, key string, values ...interface{}) error {
	return RDB.HSet(ctx, key, values...).Err()
}

// HGet 获取 Hash 字段
func HGet(ctx context.Context, key, field string) (string, error) {
	return RDB.HGet(ctx, key, field).Result()
}

// HGetAll 获取 Hash 所有字段
func HGetAll(ctx context.Context, key string) (map[string]string, error) {
	return RDB.HGetAll(ctx, key).Result()
}

// LPush 左侧插入列表
func LPush(ctx context.Context, key string, values ...interface{}) error {
	return RDB.LPush(ctx, key, values...).Err()
}

// RPop 右侧弹出列表
func RPop(ctx context.Context, key string) (string, error) {
	return RDB.RPop(ctx, key).Result()
}

// LRange 获取列表范围
func LRange(ctx context.Context, key string, start, stop int64) ([]string, error) {
	return RDB.LRange(ctx, key, start, stop).Result()
}

// SAdd 添加集合成员
func SAdd(ctx context.Context, key string, members ...interface{}) error {
	return RDB.SAdd(ctx, key, members...).Err()
}

// SMembers 获取集合所有成员
func SMembers(ctx context.Context, key string) ([]string, error) {
	return RDB.SMembers(ctx, key).Result()
}

// ZAdd 添加有序集合成员
func ZAdd(ctx context.Context, key string, members ...redis.Z) error {
	return RDB.ZAdd(ctx, key, members...).Err()
}

// ZRange 获取有序集合范围
func ZRange(ctx context.Context, key string, start, stop int64) ([]string, error) {
	return RDB.ZRange(ctx, key, start, stop).Result()
}

// Incr 自增
func Incr(ctx context.Context, key string) (int64, error) {
	return RDB.Incr(ctx, key).Result()
}

// SetNX 不存在时设置（分布式锁）
func SetNX(ctx context.Context, key string, value interface{}, expiration time.Duration) (bool, error) {
	return RDB.SetNX(ctx, key, value, expiration).Result()
}
