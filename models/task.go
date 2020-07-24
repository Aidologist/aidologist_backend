package models

import (
	"time"
	"wkBackEnd/utils/databases"
	"wkBackEnd/utils/modelsFunc"
)

type Task struct {
	Id int

	Title string
	Desc string

	Payment     int
	PayCurrency	string //enum.Currency

	PublishTime time.Time `orm:"auto_now_add;type(datetime)"`
	UpdateTime	time.Time `orm:"auto_now;type(datetime)"`
	EndTime		time.Time
	DueTime		time.Time

	// Many to One relationship
	CompanyWhoPosted 		*Company 	`orm:"rel(fk)"`    // company who posted this task
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






//--------------------------------------- CRUD Methods Below --------------------------------------
//---------------------------------- Create, Read, Update, Delete ---------------------------------
// ================================================================================================
// ================================================================================================
// ================================================================================================
// ================================================================================================
// ================================================================================================
// ================================================================================================

//--------------------------------------- Create Methods Below --------------------------------------
// Insert the Task into the database
func (this *Task) Create() {
	var _, o = modelsFunc.ConnectORM()
	_, _ = o.Insert(this)
}

//--------------------------------------- Read Methods Below --------------------------------------
// Read the Task from the database
// If the database doesn't contain this Task, the company returned will have the id of 0
func (this *Task) Read(id int) {
	this.Id = id
	var _, o = modelsFunc.ConnectORM()
	_ = o.Read(this)    // read by id of the Task
}

//--------------------------------------- Update Methods Below --------------------------------------
// Update the Task entity in the database
func (this *Task) Update() {
	var _, o = modelsFunc.ConnectORM()
	_, _ = o.Update(this)
}

//--------------------------------------- Delete Methods Below --------------------------------------
// Delete the Task by given Id
func (this *Task) Delete() {
	var _, o = modelsFunc.ConnectORM()
	_ = o.Read(this, "Id")  // find the Task by id
	_, _ = o.Delete(this)
}