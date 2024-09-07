package db

import (
	"gorm.io/gorm"
)

var (
	query *gorm.DB
)

type OrmBuilder struct{}

func NewOrmRepository() *OrmBuilder {
	return &OrmBuilder{}
}

func (orm *OrmBuilder) QueryBuilder(tx *gorm.DB) *OrmBuilder {
	query = tx
	return orm
}

func (orm *OrmBuilder) Equal(field string, value string) *OrmBuilder {
	if value != "" {
		query = query.Where(field+" = ?", value)
	}
	return orm
}

func (orm *OrmBuilder) NotEqual(field string, value string) *OrmBuilder {
	if value != "" {
		query = query.Where(field+" <> ?", value)
	}
	return orm
}

func (orm *OrmBuilder) In(field string, values []string) *OrmBuilder {
	if len(values) != 0 {
		query = query.Where(field+" IN ?", values)
	}
	return orm
}

func (orm *OrmBuilder) Likes(fields []string, value string) *OrmBuilder {
	if value != "" {
		for _, field := range fields {
			query = query.Where(field+" LIKE ?", "%"+value+"%")
		}
	}
	return orm
}

func (orm *OrmBuilder) IsBefore(field string, value string) *OrmBuilder {
	if value != "" {
		query = query.Where(field+" >= ?", value)
	}
	return orm
}

func (orm *OrmBuilder) IsBeforeLess(field string, value string) *OrmBuilder {
	if value != "" {
		query = query.Where(field+" > ?", value)
	}
	return orm
}

func (orm *OrmBuilder) IsAfter(field string, value string) *OrmBuilder {
	if value != "" {
		query = query.Where(field+" <= ?", value)
	}
	return orm
}

func (orm *OrmBuilder) IsAfterUp(field string, value string) *OrmBuilder {
	if value != "" {
		query = query.Where(field+" < ?", value)
	}
	return orm
}

func (orm *OrmBuilder) Between(field string, from string, to string) *OrmBuilder {
	if from != "" || to != "" {
		query = query.Where(field+" BETWEEM ? AND ?", from, to)
	}
	return orm
}

func (orm *OrmBuilder) Build() *gorm.DB {
	return query
}
