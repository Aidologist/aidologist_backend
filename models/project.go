package models

import "time"

type Project struct {
	Id       		int
	Title	 		string
	Desc			string
	Cover			string

	PublishTime 	time.Time	`orm:"auto_now_add;type(datetime)"`
	UpdateTime		time.Time	`orm:"auto_now;type(datetime)"`
	EndTime			time.Time
	DueTime			time.Time

	Company []*Company 			`orm:"reverse(many)"`
}