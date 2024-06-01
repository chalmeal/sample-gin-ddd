package db

// WIP
// 未完成のため利用を推奨しない

import (
	"fmt"
	"strings"
)

var (
	conditions []string
	values     []string
	query      string
)

type OrmBuilder struct{}

func NewOrmRepository() *OrmBuilder {
	return &OrmBuilder{}
}

func (orm *OrmBuilder) QueryBuilder() *OrmBuilder {
	query = ""
	conditions = []string{}
	values = []string{}
	return orm
}

func (orm *OrmBuilder) Equal(field string, value string) *OrmBuilder {
	conditions = append(conditions, fmt.Sprintf("%s = ?", field))
	values = append(values, value)
	return orm
}

func (orm *OrmBuilder) Like(field string, value string) *OrmBuilder {
	conditions = append(conditions, fmt.Sprintf("%s LIKE ?", field))
	values = append(values, "%"+value+"%")
	return orm
}

func (orm *OrmBuilder) Likes(fields []string, value string) *OrmBuilder {
	like := ""
	for _, field := range fields {
		if like != "" {
			like = like + " OR "
		}
		like = like + fmt.Sprintf("%s LIKE ?", field)
		values = append(values, "%"+value+"%")
	}
	conditions = append(conditions, like)
	return orm
}

func (orm *OrmBuilder) Between(field string) *OrmBuilder {
	conditions = append(conditions, fmt.Sprintf("%s BETWEEN ? AND ?", field))
	return orm
}

func (orm *OrmBuilder) Build() (string, []string) {
	query = strings.Join(conditions, " AND ")
	return query, values
}
