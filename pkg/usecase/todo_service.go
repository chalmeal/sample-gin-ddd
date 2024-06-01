package usecase

import (
	e "sample-gin-ddd/pkg/errors"
	"sample-gin-ddd/pkg/infrastracture/db"
	"sample-gin-ddd/pkg/infrastracture/repository"
	"sample-gin-ddd/pkg/model"
	"sample-gin-ddd/pkg/usecase/dto"
	"sample-gin-ddd/pkg/util"

	"gorm.io/gorm"
)

type TodoService struct {
	db   *gorm.DB
	rep  repository.TodoRepository
	todo model.Todos
}

func NewTodoService() *TodoService {
	return &TodoService{
		db:   db.GetDB(),
		rep:  repository.NewTodoRepository(),
		todo: *model.NewTodo(),
	}
}

// 自身のTODOを1つ取得します。
func (todo *TodoService) GetTodo(accountId string, id string) *dto.Dto {
	return db.Tx(todo.db, func(tx *gorm.DB) *dto.Dto {
		query := todo.todo.Specify(accountId, util.IdStringToInt(id))
		td, err := todo.rep.Get(tx, query)
		if err != nil {
			return &dto.Dto{
				Error: err,
			}
		}

		return &dto.Dto{
			Result: &dto.GetTodosDto{
				Todos: td,
			},
			Error: nil,
		}
	})
}

// 自身のTODO一覧を取得します。
func (todo *TodoService) FindTodos(accountId string, param *model.FindTodo) *dto.Dto {
	return db.Tx(todo.db, func(tx *gorm.DB) *dto.Dto {
		query := todo.todo.Find(accountId, param)
		// TODO: 完全一致でしか検索できていない
		td, err := todo.rep.Find(tx, query)
		if err != nil {
			return &dto.Dto{
				Error: err,
			}
		}

		return &dto.Dto{
			Result: &dto.FindTodosDto{
				Todos: td,
			},
			Error: nil,
		}
	})
}

// 自身のTODOを1つ作成します。
func (todo *TodoService) RegisterTodo(accountId string, param *model.RegTodo) *dto.Dto {
	return db.Tx(todo.db, func(tx *gorm.DB) *dto.Dto {
		reg := todo.todo.Regist(accountId, param)
		if err := todo.rep.Save(tx, reg); err != nil {
			return &dto.Dto{
				Error: err,
			}
		}

		return &dto.Dto{
			Result: &dto.RegTodoDto{
				Message: e.REGISTER_TODO_SUCCESS.Error(),
			},
		}
	})
}

// 自身のTODOを編集します。
func (todo *TodoService) EditTodo(accountId string, param *model.EdtTodo) *dto.Dto {
	return db.Tx(todo.db, func(tx *gorm.DB) *dto.Dto {
		query := todo.todo.Specify(accountId, param.TaskID)
		edt := todo.todo.Edit(accountId, param)
		if err := todo.rep.Update(tx, query, edt); err != nil {
			return &dto.Dto{
				Error: err,
			}
		}

		return &dto.Dto{
			Result: &dto.EdtTodoDto{
				Message: e.EDIT_TODO_SUCCESS.Error(),
			},
		}
	})
}

// 自身のTODOを削除します。
func (todo *TodoService) DeleteTodo(accountId string, id int) *dto.Dto {
	return db.Tx(todo.db, func(tx *gorm.DB) *dto.Dto {
		query := todo.todo.Specify(accountId, id)
		if err := todo.rep.Delete(tx, query); err != nil {
			return &dto.Dto{
				Error: err,
			}
		}

		return &dto.Dto{
			Result: &dto.DltTodoDto{
				Message: e.DELETE_TODO_SUCCESS.Error(),
			},
		}
	})
}
