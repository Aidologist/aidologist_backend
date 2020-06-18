package controllers

import (
	"github.com/astaxie/beego"
	"wkBackEnd/models"
)

type CompanyController struct {
	beego.Controller
}

// @Title Signup
// @Description
// @Success 200 {string}
// @Failure 403
// @router /signup [post]
func (u *CompanyController) Signup() {
	var company models.Company = models.Company{
		Id: 0,
		Name: u.GetString("name"),
		Password: u.GetString("password"),
		EMail: u.GetString("email")}
	company.AddCompany()
}

// @Title Login
// @Description
// @Success 200 {string}
// @Failure 403
// @router /login [post]
func (u *CompanyController) Login() {
	var company *models.Company = new(models.Company)
	company.Name = u.GetString("name")
	company.EMail = u.GetString("email")
	company.GetCompany()
	if company.Password == u.GetString("password") {
		u.Data["info"] = "success"
	} else {
		u.Data["info"] = "fail"
	}
}