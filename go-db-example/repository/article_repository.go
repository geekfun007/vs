package repository

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"time"

	"go-db-example/database"
	"go-db-example/models"

	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

type ArticleRepository struct{}

func NewArticleRepository() *ArticleRepository {
	return &ArticleRepository{}
}

// ============ CRUD ============

// Create 创建文章
func (r *ArticleRepository) Create(article *models.Article) error {
	return database.DB.Create(article).Error
}

// GetByID 根据 ID 查询（带预加载）
func (r *ArticleRepository) GetByID(ctx context.Context, id uint) (*models.Article, error) {
	cacheKey := fmt.Sprintf("article:%d", id)

	// 查缓存
	cached, err := database.Get(ctx, cacheKey)
	if err == nil && cached != "" {
		var article models.Article
		if err := json.Unmarshal([]byte(cached), &article); err == nil {
			return &article, nil
		}
	}

	// 查数据库（预加载 Author 和 Tags）
	var article models.Article
	err = database.DB.
		Preload("Author").
		Preload("Tags").
		First(&article, id).Error

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}

	// 写缓存
	if data, err := json.Marshal(article); err == nil {
		_ = database.Set(ctx, cacheKey, data, 15*time.Minute)
	}

	return &article, nil
}

// Update 更新文章
func (r *ArticleRepository) Update(ctx context.Context, article *models.Article) error {
	if err := database.DB.Save(article).Error; err != nil {
		return err
	}
	_ = database.Del(ctx, fmt.Sprintf("article:%d", article.ID))
	return nil
}

// Delete 删除文章
func (r *ArticleRepository) Delete(ctx context.Context, id uint) error {
	if err := database.DB.Delete(&models.Article{}, id).Error; err != nil {
		return err
	}
	_ = database.Del(ctx, fmt.Sprintf("article:%d", id))
	return nil
}

// ============ 复杂查询 ============

// List 分页查询文章列表
func (r *ArticleRepository) List(page, pageSize int, filters map[string]interface{}) ([]models.Article, int64, error) {
	var articles []models.Article
	var total int64

	query := database.DB.Model(&models.Article{})

	// 过滤条件
	if authorID, ok := filters["author_id"]; ok {
		query = query.Where("author_id = ?", authorID)
	}
	if status, ok := filters["status"]; ok {
		query = query.Where("status = ?", status)
	}
	if keyword, ok := filters["keyword"]; ok {
		query = query.Where("title LIKE ?", "%"+keyword.(string)+"%")
	}

	// 总数
	if err := query.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	// 分页 + 预加载
	offset := (page - 1) * pageSize
	err := query.
		Preload("Author").
		Preload("Tags").
		Offset(offset).
		Limit(pageSize).
		Order("created_at DESC").
		Find(&articles).Error

	return articles, total, err
}

// FindByAuthor 查询作者的文章
func (r *ArticleRepository) FindByAuthor(authorID uint, limit int) ([]models.Article, error) {
	var articles []models.Article
	err := database.DB.
		Where("author_id = ?", authorID).
		Order("created_at DESC").
		Limit(limit).
		Find(&articles).Error
	return articles, err
}

// FindByTag 根据标签查询文章
func (r *ArticleRepository) FindByTag(tagID uint, page, pageSize int) ([]models.Article, error) {
	var articles []models.Article
	offset := (page - 1) * pageSize

	err := database.DB.
		Joins("JOIN article_tags ON article_tags.article_id = articles.id").
		Where("article_tags.tag_id = ?", tagID).
		Preload("Author").
		Preload("Tags").
		Offset(offset).
		Limit(pageSize).
		Find(&articles).Error

	return articles, err
}

// ============ 标签操作 ============

// AddTags 添加标签
func (r *ArticleRepository) AddTags(article *models.Article, tags []models.Tag) error {
	return database.DB.Model(article).Association("Tags").Append(tags)
}

// RemoveTags 移除标签
func (r *ArticleRepository) RemoveTags(article *models.Article, tags []models.Tag) error {
	return database.DB.Model(article).Association("Tags").Delete(tags)
}

// ReplaceTags 替换标签
func (r *ArticleRepository) ReplaceTags(article *models.Article, tags []models.Tag) error {
	return database.DB.Model(article).Association("Tags").Replace(tags)
}

// ============ 浏览量（Redis 计数器） ============

// IncrViewCount 增加浏览量
func (r *ArticleRepository) IncrViewCount(ctx context.Context, articleID uint) error {
	key := fmt.Sprintf("article:views:%d", articleID)
	_, err := database.Incr(ctx, key)
	return err
}

// GetViewCount 获取浏览量
func (r *ArticleRepository) GetViewCount(ctx context.Context, articleID uint) (int64, error) {
	key := fmt.Sprintf("article:views:%d", articleID)
	val, err := database.RDB.Get(ctx, key).Int64()
	if errors.Is(err, redis.Nil) {
		return 0, nil
	}
	return val, err
}

// SyncViewCountToDB 同步浏览量到数据库（定时任务调用）
func (r *ArticleRepository) SyncViewCountToDB(ctx context.Context, articleID uint) error {
	key := fmt.Sprintf("article:views:%d", articleID)
	count, err := database.RDB.GetDel(ctx, key).Int64()
	if err != nil && !errors.Is(err, redis.Nil) {
		return err
	}
	if count > 0 {
		return database.DB.Model(&models.Article{}).
			Where("id = ?", articleID).
			UpdateColumn("view_count", gorm.Expr("view_count + ?", count)).Error
	}
	return nil
}

// ============ 热门文章（Redis 有序集合） ============

// AddToHotRank 添加到热门排行
func (r *ArticleRepository) AddToHotRank(ctx context.Context, articleID uint, score float64) error {
	return database.ZAdd(ctx, "article:hot", redis.Z{
		Score:  score,
		Member: articleID,
	})
}

// GetHotArticles 获取热门文章 ID
func (r *ArticleRepository) GetHotArticles(ctx context.Context, limit int) ([]string, error) {
	return database.RDB.ZRevRange(ctx, "article:hot", 0, int64(limit-1)).Result()
}
