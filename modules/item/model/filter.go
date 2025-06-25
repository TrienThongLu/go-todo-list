package model

type Filter struct {
	Status ToDoStatus `json:"status" form:"status"`
}
