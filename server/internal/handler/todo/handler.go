package todo

import (
	"github.com/mohwa/ci-cd-github-action/internal/db"
)

func GetTodos() ([]*db.Todo, error) {
	return db.GetTodos()
}

func CreateTodo(todo *db.Todo) error {
	return db.CreateTodo(todo)
}
