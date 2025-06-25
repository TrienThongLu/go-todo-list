package model

import (
	"errors"

	"github.com/TrienThongLu/GoREST/common"
)

var (
	ErrTitleIsBlank = errors.New("title cannot be blank")
	ErrItemDeleted  = errors.New("item is deleted")
)

type ToDoItem struct {
	common.SQLModel
	Title       string     `json:"title" gorm:"column:title"`
	Description string     `json:"description" gorm:"column:description"`
	Status      ToDoStatus `json:"status" gorm:"column:status"`
}

func (ToDoItem) TableName() string { return "todo_items" }

type ToDoItemCreation struct {
	Id          int    `json:"-" gorm:"column:id"`
	Title       string `json:"title" gorm:"column:title"`
	Description string `json:"description" gorm:"column:description"`
}

func (ToDoItemCreation) TableName() string { return ToDoItem{}.TableName() }

type ToDoItemUpdate struct {
	Title       *string     `json:"title" gorm:"column:title"`
	Description string      `json:"description" gorm:"column:description"`
	Status      *ToDoStatus `json:"status" gorm:"column:status"`
}

func (ToDoItemUpdate) TableName() string { return ToDoItem{}.TableName() }
