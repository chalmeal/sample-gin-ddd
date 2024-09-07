package repository

import (
	"log"
	e "sample-gin-ddd/pkg/errors"
	"sample-gin-ddd/pkg/model"
	"sample-gin-ddd/pkg/util"

	"gorm.io/gorm"
)

type TodoRepository interface {
	Get(tx *gorm.DB, query interface{}) (*model.Todos, error)
	Find(query *gorm.DB) (*[]model.Todos, error)
	Save(tx *gorm.DB, param *model.Todos) error
	Update(tx *gorm.DB, query interface{}, param *model.Todos) error
	Delete(tx *gorm.DB, query interface{}) error
}

type TodoRepositoryImpl struct{}

func NewTodoRepository() TodoRepository {
	return &TodoRepositoryImpl{}
}

func (rep *TodoRepositoryImpl) Get(tx *gorm.DB, query interface{}) (*model.Todos, error) {
	var todo model.Todos
	err := tx.Where(query).First(&todo)
	if err.Error != nil {
		return &model.Todos{}, e.GET_TODO_NOT_FOUND
	}
	return &todo, nil
}

func (rep *TodoRepositoryImpl) Find(query *gorm.DB) (*[]model.Todos, error) {
	var todo []model.Todos
	err := query.Find(&todo)
	if err.Error != nil {
		log.Println(err)
		return &[]model.Todos{}, e.INTERNAL_SERVER_ERROR
	}
	return &todo, nil
}

func (rep *TodoRepositoryImpl) Save(tx *gorm.DB, param *model.Todos) error {
	if err := tx.Create(param).Error; err != nil {
		log.Println(err)
		return e.INTERNAL_SERVER_ERROR
	}
	return nil
}

func (rep *TodoRepositoryImpl) Update(tx *gorm.DB, query interface{}, param *model.Todos) error {
	var todo model.Todos
	record := tx.Where(query).First(&todo)
	if record.Error != nil || util.IsEmpty(todo, model.Todos{}) {
		log.Println(record.Error)
		return e.GET_TODO_NOT_FOUND
	}
	if err := record.Updates(param).Error; err != nil {
		log.Println(err)
		return e.INTERNAL_SERVER_ERROR
	}
	return nil
}

func (rep *TodoRepositoryImpl) Delete(tx *gorm.DB, query interface{}) error {
	var todo model.Todos
	record := tx.Where(query).First(&todo)
	if record.Error != nil || util.IsEmpty(todo, model.Todos{}) {
		log.Println(record.Error)
		return e.GET_TODO_NOT_FOUND
	}
	if err := record.Delete(&todo).Error; err != nil {
		log.Println(err)
		return e.INTERNAL_SERVER_ERROR
	}
	return nil
}
