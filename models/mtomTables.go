package models

import (
	"time"
	"wkBackEnd/utils/modelsFunc"
)
// This file contains all the third tables that many to many relationship yield
// We can manges these third tables using data models, then we can add different attributes in the third tables

//--------------------------------------- Data Models Below ---------------------------------------
// ================================================================================================
// ================================================================================================
// ================================================================================================
// ================================================================================================
// ================================================================================================
// ================================================================================================

type CompanyFavoriteUser struct {
	Id    int64
	Company    *Company    `orm:"rel(fk)"`
	User    *User    `orm:"rel(fk)"`
	FavoriteAt   time.Time    `orm:"auto_now_add;type(datetime)"`
}

func (CompanyFavoriteUser) TableName() string{
	return "company_favorite_user"
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
// Insert the CompanyFavoriteUser into the database
func (this *CompanyFavoriteUser) Create() {
	var _, o = modelsFunc.ConnectORM()
	_, _ = o.Insert(this)
}

//--------------------------------------- Read Methods Below --------------------------------------
// Read the CompanyFavoriteUser from the database
// If the database doesn't contain this CompanyFavoriteUser, the CompanyFavoriteUser returned will have the id of 0
func (this *CompanyFavoriteUser) Read() {
	var _, o = modelsFunc.ConnectORM()
	_ = o.Read(this)    // read by id of the CompanyFavoriateUser
}

//--------------------------------------- Update Methods Below --------------------------------------
// Update the CompanyFavoriteUser entity in the database
func (this *CompanyFavoriteUser) Update() {
	var _, o = modelsFunc.ConnectORM()
	_, _ = o.Update(this)
}

//--------------------------------------- Delete Methods Below --------------------------------------
// Delete the CompanyFavoriteUser by given Id
func (this *CompanyFavoriteUser) Delete() {
	var _, o = modelsFunc.ConnectORM()
	_ = o.Read(this, "Id")  // find the CompanyFavoriteUser by id
	_, _ = o.Delete(this)
}