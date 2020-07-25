package models

import (
	"time"
	"wkBackEnd/models"
)

type UserInChatRoom struct {
	Id    						int
	NickName					string

	ChatRoom    *ChatRoom		`orm:"rel(fk)"`
	User    	*models.User	`orm:"rel(fk)"`
	JoinAt  	time.Time   	`orm:"auto_now_add;type(datetime)"`
}