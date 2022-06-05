package db

import (
	"fmt"

	"gorm.io/gorm"
)

const tableName = "todo"

func TodoTable() (*gorm.DB, error) {
	if !isExistTable(tableName) {
		return nil, fmt.Errorf("%w: %s", ErrorNoTable, tableName)
	}

	return db.Table(tableName).Session(&gorm.Session{}), nil
}

type Todo struct {
	ID   int    `gorm:"primaryKey"`
	Name string `json:"name" gorm:"type:varchar(256); column:name"`
}

func (Todo) TableName() string {
	return tableName
}

func GetTodos() ([]*Todo, error) {
	table, err := TodoTable()
	if err != nil {
		return nil, ErrorTableOpenFail
	}

	var ret []*Todo
	// &ret Todo 구조체 포인터
	result := table.Find(&ret)

	if result.Error != nil {
		return nil, result.Error
	}

	return ret, nil
}

func CreateTodo(todo *Todo) error {
	table, err := TodoTable()

	if err != nil {
		return ErrorTableOpenFail
	}

	result := table.Create(&todo)

	if result.Error != nil {
		return result.Error
	}

	return nil
}
