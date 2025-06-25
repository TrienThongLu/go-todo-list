package model

type ToDoStatus string

const (
	StatusDoing   ToDoStatus = "Doing"
	StatusDone    ToDoStatus = "Done"
	StatusDeleted ToDoStatus = "Deleted"
)
