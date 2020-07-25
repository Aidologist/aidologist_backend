package models

import (
	"time"
	models "wkBackEnd/models/chat"
	"wkBackEnd/utils/databases"
	"wkBackEnd/utils/modelsFunc"
)

//--------------------------------------- Data Models Below ---------------------------------------
// ================================================================================================
// ================================================================================================
// ================================================================================================
// ================================================================================================
// ================================================================================================
// ================================================================================================

type User struct {
	Id       int
	EMail 	 string
	Username string
	Password string
	CreateTime 	time.Time	`orm:"auto_now_add;type(datetime)"`
	UpdateTime  time.Time	`orm:"auto_now;type(datetime)"`

	//Profile  *Profile `orm:"null;rel(one);on_delete(set_null)"`

	// One to One
	SendMessages	[]*models.Message `orm:"reverse(many)"`

	// One to Many

	// Many to many
	FavoriteTasks  []*Task `orm:"rel(m2m);rel_through(wkBackEnd/models.UserFavoriteTask)"`     // Many to Many with User
	ChatRoom 		[]*models.ChatRoom `orm:"rel(m2m);rel_through(wkBackEnd/models.UserInChatRoom)"`

	// Reverse relationship
	CompaniesWhoFavorite   []*Company    `orm:"reverse(many)"` // Reverse of Many to Many with CompanyUser, users liked by the CompanyUser
}

//------------------------------------ Attached Data model below ----------------------------------
// ================================================================================================

type Supervisor struct {

}

type CompanyUser struct {

}

//-------------------------------------- Unused Methods Below -------------------------------------
//--------------------------------------- But not yet deleted -------------------------------------
// ================================================================================================

func (this *User) AddUser() (int64, bool) {
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

func (this *User) GetUser() bool {
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


//--------------------------------------- CRUD Methods Below --------------------------------------
//---------------------------------- Create, Read, Update, Delete ---------------------------------
// ================================================================================================
// ================================================================================================
// ================================================================================================
// ================================================================================================
// ================================================================================================
// ================================================================================================

//--------------------------------------- Create Methods Below --------------------------------------
// Insert the User into the database
// TODO: Check if the user already existed in the database
func (this *User) Create() {
	var _, o = modelsFunc.ConnectORM()
	_, _ = o.Insert(this)
}


//--------------------------------------- Read Methods Below --------------------------------------
// Read the User from the database
// TODO: Check if the user already existed in the database
func (this *User) Read() {
	var _, o = modelsFunc.ConnectORM()
	_ = o.Read(this)
}

//--------------------------------------- Update Methods Below --------------------------------------
// Update the User entity
// TODO: Check if the user already existed in the database
func (this *User) Update() {
	var _, o = modelsFunc.ConnectORM()
	_, _ = o.Update(this)
}


//--------------------------------------- Delete Methods Below --------------------------------------
// TODO: Check if the user already existed in the database
// Delete the User by given Id
func (this *User) Delete() {
	var _, o = modelsFunc.ConnectORM()
	_ = o.Read(this, "Id")  // find the User by id
	_, _ = o.Delete(this)
}

