package dto

import "sample-gin-ddd/pkg/model"

type GetTodosDto struct {
	Todos *model.Todos
}

type FindTodosDto struct {
	Todos *[]model.Todos
}

type RegTodoDto struct {
	Message string
}

type EdtTodoDto struct {
	Message string
}

type DltTodoDto struct {
	Message string
}
