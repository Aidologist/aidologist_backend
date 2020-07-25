package models

import (
	"time"
	"wkBackEnd/utils/modelsFunc"
)

type ChatRoom struct {
	Id       	int
	Name 		string
	Desc		string
	CreateTime 	time.Time		`orm:"auto_now_add;type(datetime)"`
	UpdateTime  time.Time		`orm:"auto_now;type(datetime)"`

	Users		[]*User   `orm:"reverse(many)"`
	Messages	[]*Message `orm:"reverse(many)"`
}

type UserInChatRoom struct {
	Id    						int
	NickName					string

	ChatRoom    *ChatRoom `orm:"rel(fk)"`
	User    	*User     `orm:"rel(fk)"`
	JoinAt  	time.Time   `orm:"auto_now_add;type(datetime)"`
}

//--------------------------------------- CRUD Methods Below --------------------------------------
//---------------------------------- Create, Read, Update, Delete ---------------------------------
// ================================================================================================

//--------------------------------------- Create Methods Below --------------------------------------
func (this *ChatRoom) Create() {
	var _, o = modelsFunc.ConnectORM()
	_, _ = o.Insert(this)
}

//--------------------------------------- Read Methods Below --------------------------------------
func (this *ChatRoom) Read() {
	var _, o = modelsFunc.ConnectORM()
	_ = o.Read(this)
}

//--------------------------------------- Update Methods Below --------------------------------------
func (this *ChatRoom) Update() {
	var _, o = modelsFunc.ConnectORM()
	_, _ = o.Update(this)
}

//--------------------------------------- Delete Methods Below --------------------------------------
func (this *ChatRoom) Delete() {
	var _, o = modelsFunc.ConnectORM()
	_ = o.Read(this, "Id")
	_, _ = o.Delete(this)
}

//--------------------------------------- Create Methods Below --------------------------------------
func (this *UserInChatRoom) Create() {
	var _, o = modelsFunc.ConnectORM()
	_, _ = o.Insert(this)
}

//--------------------------------------- Read Methods Below --------------------------------------
func (this *UserInChatRoom) Read() {
	var _, o = modelsFunc.ConnectORM()
	_ = o.Read(this)
}

//--------------------------------------- Update Methods Below --------------------------------------
func (this *UserInChatRoom) Update() {
	var _, o = modelsFunc.ConnectORM()
	_, _ = o.Update(this)
}

//--------------------------------------- Delete Methods Below --------------------------------------
func (this *UserInChatRoom) Delete() {
	var _, o = modelsFunc.ConnectORM()
	_ = o.Read(this, "Id")
	_, _ = o.Delete(this)
}