package models

type NotificationTemplate struct {
	Id       int
	Name     string
	Type     string
	Content  string
	IsActive bool
}
