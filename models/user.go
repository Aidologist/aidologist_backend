package models

import (
	"time"
	"wkBackEnd/utils/databases"
)

type User struct {
	Id       		int
	EMail 	 		string
	Username 		string
	Password 		string
	CreateTime 		time.Time	`orm:"auto_now_add;type(datetime)"`
	LastLoginTime  	time.Time

	// One to One

	// One to Many

	// Many to many

	// Reverse relationships
}

// ================================================================================================
//--------------------------------------- CRUD Methods Below --------------------------------------
//---------------------------------- Create, Read, Update, Delete ---------------------------------
// ================================================================================================

//--------------------------------------- Create Methods Below ------------------------------------
func (this *User) Create() {
	var _, o = databases.ConnectOrm()
	_, _ = o.Insert(this)
}

//--------------------------------------- Read Methods Below --------------------------------------
func (this *User) Read() {
	var _, o = databases.ConnectOrm()
	_ = o.Read(this)
}

func (this *User) ReadByEmail() {
	var _, o = databases.ConnectOrm()
	_ = o.Read(this, "e_mail")
}

//--------------------------------------- Update Methods Below ------------------------------------
func (this *User) Update() {
	var _, o = databases.ConnectOrm()
	_, _ = o.Update(this)
}

//--------------------------------------- Delete Methods Below ------------------------------------
func (this *User) Delete() {
	var _, o = databases.ConnectOrm()
	_ = o.Read(this, "Id")  // find the User by id
	_, _ = o.Delete(this)
}

// ================================================================================================
//------------------------------------- Advanced Methods Below ------------------------------------
//------------------------------ Advanced CRUD, special model methods -----------------------------
// ================================================================================================
