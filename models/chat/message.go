package models

import (
	"time"
	"wkBackEnd/models"
	"wkBackEnd/utils/constants/enum"
)

type Message struct {
	Id			int
	Type		enum.MessageType
	Contain 	string
	Time 		time.Time		`orm:"auto_now_add;type(datetime)"`

	User		*models.User	`orm:"rel(fk)"`
	InRoom		*ChatRoom		`orm:"rel(fk)"`
}