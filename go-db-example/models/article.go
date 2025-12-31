package models

import (
	"time"

	"gorm.io/gorm"
)

// Article 文章模型
type Article struct {
	ID        uint           `gorm:"primaryKey" json:"id"`
	Title     string         `gorm:"size:200;not null" json:"title"`
	Content   string         `gorm:"type:text" json:"content"`
	Summary   string         `gorm:"size:500" json:"summary"`
	AuthorID  uint           `gorm:"index;not null" json:"author_id"`
	Author    User           `gorm:"foreignKey:AuthorID" json:"author,omitempty"`
	Status    int8           `gorm:"default:0;comment:0=draft,1=published" json:"status"`
	ViewCount int            `gorm:"default:0" json:"view_count"`
	Tags      []Tag          `gorm:"many2many:article_tags" json:"tags,omitempty"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`
}

// Tag 标签模型
type Tag struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	Name      string    `gorm:"size:50;uniqueIndex;not null" json:"name"`
	Articles  []Article `gorm:"many2many:article_tags" json:"articles,omitempty"`
	CreatedAt time.Time `json:"created_at"`
}

func (Article) TableName() string {
	return "articles"
}

func (Tag) TableName() string {
	return "tags"
}
