package models

import "time"

type Article struct {
	Id       		int
	Title	 		string
	Abstract		string
	Source			string

	PublishTime 	time.Time	`orm:"auto_now_add;type(datetime)"`
	UpdateTime		time.Time	`orm:"auto_now;type(datetime)"`

	User			*[]User
	Viewed			*[]User
}