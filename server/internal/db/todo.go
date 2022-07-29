package db

import (
	"fmt"

	"gorm.io/gorm"
)

const todoTableName = "todo"

func TodoTable() (*gorm.DB, error) {
	if !isExistTable(todoTableName) {
		return nil, fmt.Errorf("%w: %s", ErrorNoTable, todoTableName)
	}

	return db.Table(todoTableName).Session(&gorm.Session{}), nil
}

type Todo struct {
	ID   int    `gorm:"primaryKey"`
	Name string `gorm:"type:varchar(256); column:name"`
}

func (Todo) TableName() string {
	return todoTableName
}

func GetTodos() ([]*Todo, error) {
	table, err := TodoTable()
	if err != nil {
		return nil, ErrorTableOpenFail
	}

	var ret []*Todo
	// 참조값(&ret)을 전달 후, Find 함수를 통해, 그 원본 값(ret)을 수정하게한다.
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
