package model

import (
	"sample-gin-ddd/pkg/util"
	"time"
)

type Todos struct {
	TaskID    uint       `gorm:"primaryKey; autoIncrement" json:"task_id"`
	AccountID string     `gorm:"type:varchar(50); not null" json:"account_id"`
	Title     string     `gorm:"type:varchar(100); not null" json:"title"`
	Detail    string     `gorm:"type:varchar(1000)" json:"detail"`
	Category  string     `gorm:"type:varchar(20)" json:"category"`
	Status    string     `gorm:"type:varchar(20); not null" json:"status"`
	ExpiredAt *time.Time `gorm:"type:datetime" json:"expired_at"`
	CreatedAt time.Time  `gorm:"type:datetime; not null" json:"created_at"`
	UpdatedAt *time.Time `gorm:"type:datetime; autoUpdateTime:false" json:"updated_at"`
	DeletedAt *time.Time `gorm:"type:datetime" json:"deleted_at"`
}

func NewTodo() *Todos {
	return &Todos{}
}

type SpecifyTodo struct {
	TaskID int `json:"task_id"`
}

func (todo *Todos) Specify(accountId string, id int) *Todos {
	return &Todos{
		TaskID:    uint(id),
		AccountID: accountId,
	}
}

type FindTodo struct {
	Title    string `json:"title"`
	Category string `json:"category"`
	Status   string `json:"status"`
}

func (todo *Todos) Find(accountId string, p *FindTodo) *Todos {
	return &Todos{
		AccountID: accountId,
		Title:     p.Title,
		Category:  p.Category,
		Status:    p.Status,
	}
}

type RegTodo struct {
	Title     string `json:"title" binding:"required"`
	Detail    string `json:"detail"`
	Category  string `json:"category"`
	Status    string `json:"status" binding:"required"`
	ExpiredAt string `json:"expired_at"`
}

func (todo *Todos) Regist(accountId string, p *RegTodo) *Todos {
	return &Todos{
		AccountID: accountId,
		Title:     p.Title,
		Detail:    p.Detail,
		Category:  p.Category,
		Status:    p.Status,
		ExpiredAt: util.ParseStringTime(p.ExpiredAt),
	}
}

type EdtTodo struct {
	TaskID    int    `json:"task_id" binding:"required"`
	Title     string `json:"title" binding:"required"`
	Detail    string `json:"detail"`
	Category  string `json:"category"`
	Status    string `json:"status" binding:"required"`
	ExpiredAt string `json:"expired_at"`
}

func (todo *Todos) Edit(accountId string, p *EdtTodo) *Todos {
	return &Todos{
		Title:     p.Title,
		Detail:    p.Detail,
		Category:  p.Category,
		Status:    p.Status,
		ExpiredAt: util.ParseStringTime(p.ExpiredAt),
		UpdatedAt: util.NowDateTime(),
	}
}
