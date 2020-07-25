package models

import (
	"time"
	"wkBackEnd/models"
)

type ChatRoom struct {
	Id       	int
	Name 		string
	Desc		string
	CreateTime 	time.Time		`orm:"auto_now_add;type(datetime)"`
	UpdateTime  time.Time		`orm:"auto_now;type(datetime)"`

	Users		[]*models.User	`orm:"reverse(many)"`
	Messages	[]*Message		`orm:"reverse(many)"`
}