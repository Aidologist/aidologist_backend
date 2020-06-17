package models

import (
	"time"
	"wkBackEnd/utils/databases"
)

type Task struct {
	Id int

	Title string
	Desc string

	Payment     int
	PayCurrency	string //enum.Currency

	PublishTIme time.Time
	UpdateTime	time.Time
	EndTime		time.Time
	DueTime		time.Time
}

func (this *Task) AddTask() (int64, bool) {
	success, o := databases.ConnectOrm()
	if success {
		index, err := o.Insert(this)
		if err == nil {
			return index, true
		} else {
			return 0, false
		}
	} else {
		return 0, false
	}
}

func (this *Task) UpdateTask() (int64, bool) {
	success, o := databases.ConnectOrm()
	if success {
		index, err := o.Update(this)
		if err == nil {
			return index, true
		} else {
			return 0, false
		}
	} else {
		return 0, false
	}
}

func (this *Task) ReadTaskbyId(id int) bool {
	this.Id = id
	success, o := databases.ConnectOrm()
	if success {
		err := o.Read(this)
		if err == nil {
			return true
		} else {
			return false
		}
	} else {
		return false
	}
}

func (this *Task) ReadTask() bool {
	success, o := databases.ConnectOrm()
	if success {
		err := o.QueryTable(this).One(this)
		if err == nil {
			return true
		} else {
			return false
		}
	} else {
		return false
	}
}