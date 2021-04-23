package model

import (
	"time"
)

// ListParams 列表参数
type ListParams struct {
	GetCount bool //是否统计总数（不含Limit和Offset的记录集总数）
	Limit    int
	Offset   int
	Page     int
}

type User struct {
	ID        int
	Name      string
	CreatedAt time.Time
	UpdatedAt time.Time
	Articles  []Article `gorm:"foreignKey:AuthorId"`
}

type Article struct {
	ID         int
	Title      string
	Content    string
	AuthorId   int `sql:"index"`
	AuthorName string
	CreatedAt  time.Time
	UpdatedAt  time.Time
}
