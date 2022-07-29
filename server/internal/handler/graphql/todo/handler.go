package todo

import (
	"github.com/mohwa/ci-cd-github-action/api/graphql/model"
	"github.com/mohwa/ci-cd-github-action/internal/db"
)

func GetTodos() ([]*db.Todo, error) {
	return db.GetTodos()
}

func CreateTodo(todo model.TodoInput) error {
	return db.CreateTodo(&db.Todo{Name: todo.Name})
}
