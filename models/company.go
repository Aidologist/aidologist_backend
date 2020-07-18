package models

import (
	"time"
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

type Company struct {
	Id       int
	EMail 	 string
	Name	 string
	Password string
	CreateTime 	time.Time	`orm:"auto_now_add;type(datetime)"`
	UpdateTime  time.Time	`orm:"auto_now;type(datetime)"`

	// One to One

	// One to Many
	Projects []*Project `orm:"reverse(many)"`

	// Many to many
	FavoriteUsers  []*User `orm:"rel(m2m);rel_through(wkBackEnd/models.CompanyFavoriteUser)"`     // Many to Many with User
	// Reverse relationship

}

func (this *Company) AddCompany() (int64, bool) {
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

func (this *Company) GetCompany() bool {
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
// Insert the Company into the database
func (this *Company) Create() {
	var _, o = modelsFunc.ConnectORM()
	_, _ = o.Insert(this)
}

//--------------------------------------- Read Methods Below --------------------------------------
// Read the Company from the database
// If the database doesn't contain this company, the company returned will have the id of 0
func (this *Company) Read(id int) {
	this.Id = id
	var _, o = modelsFunc.ConnectORM()
	_ = o.Read(this)    // read by id of the company
}

//--------------------------------------- Update Methods Below --------------------------------------
// Update the Company entity in the database
func (this *Company) Update() {
	var _, o = modelsFunc.ConnectORM()
	_, _ = o.Update(this)
}

//--------------------------------------- Delete Methods Below --------------------------------------
// Delete the Company by given Id
func (this *Company) Delete() {
	var _, o = modelsFunc.ConnectORM()
	_ = o.Read(this, "Id")  // find the Company by id
	_, _ = o.Delete(this)
}

//--------------------------------------- Advanced Method Below --------------------------------------
func (this *Company) ReadAllProject() {
	var _, o = modelsFunc.ConnectORM()
	o.QueryTable("Project").Filter("Company__Id", this.Id).All(&this.Projects)
// .Filter("Company", this.Id)
}