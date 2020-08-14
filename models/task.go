package models

import "time"

type Task struct {
	Id       			int
	Title 	 			string
	TaskDesc			string
	TaskRequirement 	string
	TaskType			int
	PayMethod			int
	PayCurrency			int
	Payment				int
	CreateTime 			time.Time	`orm:"auto_now_add;type(datetime)"`
	UpdateTime  		time.Time	`orm:"auto_now;type(datetime)"`
	PublishTime			time.Time
	ClosedTime			time.Time
	DueTime				time.Time

	// ---------- ApplyTask fields ----------

	// ---------- BidTask fields ----------

	// ---------- QuickTask fields ----------
	//MaxAccept 			int
	//CurrentAccept 		int

	// ========== One to One ==========

	// ========== One to Many ==========

	// ========== Many to many ==========

	// ========== Reverse relationships ==========
}