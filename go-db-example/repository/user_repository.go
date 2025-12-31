package repository

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"time"

	"go-db-example/database"
	"go-db-example/models"

	"gorm.io/gorm"
)

type UserRepository struct{}

func NewUserRepository() *UserRepository {
	return &UserRepository{}
}

// ============ GORM CRUD 操作 ============

// Create 创建用户
func (r *UserRepository) Create(user *models.User) error {
	return database.DB.Create(user).Error
}

// GetByID 根据 ID 查询（带缓存）
func (r *UserRepository) GetByID(ctx context.Context, id uint) (*models.User, error) {
	cacheKey := fmt.Sprintf("user:%d", id)

	// 1. 先查 Redis 缓存
	cached, err := database.Get(ctx, cacheKey)
	if err == nil && cached != "" {
		var user models.User
		if err := json.Unmarshal([]byte(cached), &user); err == nil {
			return &user, nil
		}
	}

	// 2. 缓存未命中，查 MySQL
	var user models.User
	if err := database.DB.First(&user, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}

	// 3. 写入缓存
	if data, err := json.Marshal(user); err == nil {
		_ = database.Set(ctx, cacheKey, data, 30*time.Minute)
	}

	return &user, nil
}

// GetByUsername 根据用户名查询
func (r *UserRepository) GetByUsername(username string) (*models.User, error) {
	var user models.User
	err := database.DB.Where("username = ?", username).First(&user).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	}
	return &user, err
}

// GetByEmail 根据邮箱查询
func (r *UserRepository) GetByEmail(email string) (*models.User, error) {
	var user models.User
	err := database.DB.Where("email = ?", email).First(&user).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	}
	return &user, err
}

// Update 更新用户（清除缓存）
func (r *UserRepository) Update(ctx context.Context, user *models.User) error {
	if err := database.DB.Save(user).Error; err != nil {
		return err
	}
	// 清除缓存
	cacheKey := fmt.Sprintf("user:%d", user.ID)
	_ = database.Del(ctx, cacheKey)
	return nil
}

// UpdateFields 更新指定字段
func (r *UserRepository) UpdateFields(ctx context.Context, id uint, fields map[string]interface{}) error {
	if err := database.DB.Model(&models.User{}).Where("id = ?", id).Updates(fields).Error; err != nil {
		return err
	}
	cacheKey := fmt.Sprintf("user:%d", id)
	_ = database.Del(ctx, cacheKey)
	return nil
}

// Delete 软删除用户
func (r *UserRepository) Delete(ctx context.Context, id uint) error {
	if err := database.DB.Delete(&models.User{}, id).Error; err != nil {
		return err
	}
	cacheKey := fmt.Sprintf("user:%d", id)
	_ = database.Del(ctx, cacheKey)
	return nil
}

// HardDelete 硬删除
func (r *UserRepository) HardDelete(ctx context.Context, id uint) error {
	if err := database.DB.Unscoped().Delete(&models.User{}, id).Error; err != nil {
		return err
	}
	cacheKey := fmt.Sprintf("user:%d", id)
	_ = database.Del(ctx, cacheKey)
	return nil
}

// ============ 查询方法 ============

// List 分页查询
func (r *UserRepository) List(page, pageSize int, conditions map[string]interface{}) ([]models.User, int64, error) {
	var users []models.User
	var total int64

	query := database.DB.Model(&models.User{})

	// 动态条件
	if status, ok := conditions["status"]; ok {
		query = query.Where("status = ?", status)
	}
	if keyword, ok := conditions["keyword"]; ok {
		query = query.Where("username LIKE ? OR nickname LIKE ?",
			"%"+keyword.(string)+"%", "%"+keyword.(string)+"%")
	}

	// 统计总数
	if err := query.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	// 分页查询
	offset := (page - 1) * pageSize
	if err := query.Offset(offset).Limit(pageSize).Order("id DESC").Find(&users).Error; err != nil {
		return nil, 0, err
	}

	return users, total, nil
}

// FindByIDs 批量查询
func (r *UserRepository) FindByIDs(ids []uint) ([]models.User, error) {
	var users []models.User
	if err := database.DB.Where("id IN ?", ids).Find(&users).Error; err != nil {
		return nil, err
	}
	return users, nil
}

// Exists 检查是否存在
func (r *UserRepository) Exists(field, value string) (bool, error) {
	var count int64
	err := database.DB.Model(&models.User{}).Where(field+" = ?", value).Count(&count).Error
	return count > 0, err
}

// ============ 事务示例 ============

// CreateWithProfile 事务：创建用户同时创建关联数据
func (r *UserRepository) CreateWithProfile(user *models.User) error {
	return database.DB.Transaction(func(tx *gorm.DB) error {
		// 创建用户
		if err := tx.Create(user).Error; err != nil {
			return err
		}

		// 可以在这里创建其他关联数据
		// if err := tx.Create(&profile).Error; err != nil {
		//     return err  // 返回错误会自动回滚
		// }

		return nil // 返回 nil 提交事务
	})
}
