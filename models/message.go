package models

import (
	"time"
	"wkBackEnd/utils/constants/enum"
	"wkBackEnd/utils/modelsFunc"
)

type Message struct {
	Id			int
	Type		enum.MessageType
	Contain 	string
	Time 		time.Time		`orm:"auto_now_add;type(datetime)"`

	User		*User             `orm:"rel(fk)"`
	InRoom		*ChatRoom       `orm:"rel(fk)"`
	Tracks		[]*MessageTrack `orm:"rel(m2m);rel_through(wkBackEnd/models.MessageInTrack)"`
}

type MessageTrack struct {
	Id			int
	Name		string
	Desc		string
	CreateTime 	time.Time		`orm:"auto_now_add;type(datetime)"`
	UpdateTime  time.Time		`orm:"auto_now;type(datetime)"`

	Messages	[]*Message            `orm:"reverse(many)"`
	Status		*MessageTrackStatus `orm:"rel(fk)"`
}

type MessageTrackStatus struct {
	Id			int
	Name		string

	Track 		[]*MessageTrack `orm:"reverse(many)"`
}

type MessageInTrack struct {
	Id				int
	Message			*Message    `orm:"rel(fk)"`
	Track    		*MessageTrack `orm:"rel(fk)"`
}

//--------------------------------------- CRUD Methods Below --------------------------------------
//---------------------------------- Create, Read, Update, Delete ---------------------------------
// ================================================================================================

//--------------------------------------- Create Methods Below --------------------------------------
func (this *Message) Create() {
	var _, o = modelsFunc.ConnectORM()
	_, _ = o.Insert(this)
}

//--------------------------------------- Read Methods Below --------------------------------------
func (this *Message) Read() {
	var _, o = modelsFunc.ConnectORM()
	_ = o.Read(this)
}

//--------------------------------------- Update Methods Below --------------------------------------
func (this *Message) Update() {
	var _, o = modelsFunc.ConnectORM()
	_, _ = o.Update(this)
}

//--------------------------------------- Delete Methods Below --------------------------------------
func (this *Message) Delete() {
	var _, o = modelsFunc.ConnectORM()
	_ = o.Read(this, "Id")
	_, _ = o.Delete(this)
}