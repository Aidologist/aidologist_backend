package models

import (
	"time"
	"wkBackEnd/utils/modelsFunc"
)

type Project struct {
	Id       		int
	Title	 		string
	Desc			string
	Cover			string

	PublishTime 	time.Time	`orm:"auto_now_add;type(datetime)"`
	UpdateTime		time.Time	`orm:"auto_now;type(datetime)"`
	EndTime			time.Time
	DueTime			time.Time

	// Many to One relationship
	Company 		*Company 	`orm:"rel(fk)"`
}

//--------------------------------------- Create Methods Below --------------------------------------
// Insert the User into the database
// TODO: Check if the user already existed in the database
func (this *Project) Create() {
	var _, o = modelsFunc.ConnectORM()
	_, _ = o.Insert(this)
}


//--------------------------------------- Read Methods Below --------------------------------------
// Read the User from the database
// TODO: Check if the user already existed in the database
func (this *Project) Read(id int) {
	this.Id = id
	var _, o = modelsFunc.ConnectORM()
	_ = o.Read(this)
}

//--------------------------------------- Update Methods Below --------------------------------------
// Update the User entity
// TODO: Check if the user already existed in the database
func (this *Project) Update() {
	var _, o = modelsFunc.ConnectORM()
	_, _ = o.Update(this)
}


//--------------------------------------- Delete Methods Below --------------------------------------
// TODO: Check if the user already existed in the database
// Delete the User by given Id
func (this *Project) Delete() {
	var _, o = modelsFunc.ConnectORM()
	_ = o.Read(this, "Id")  // find the User by id
	_, _ = o.Delete(this)
}