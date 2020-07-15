package controllers

import (
	"github.com/astaxie/beego"
	"strconv"
	"wkBackEnd/models"
)

type CompanyController struct {
	beego.Controller
}

// tested and successfully created a user in the database
// the id also atomatically increased
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
	//company.AddCompany
	company.Create()
}

// tested and successfully delete a company in the database
// @Title DeleteCompany
// @Description Delete company by company ID
// @Success 200 {string}
// @Failure 403
// @router /deleteCompany [post]
func (u *CompanyController) DeleteCompany() {
	companyId := u.GetString("companyID")
	companyid,_ := strconv.Atoi(companyId)
	var company models.Company = models.Company{
		Id: companyid}
	company.Delete()
}

// Tested sucessfully
// But you can not leave companyname,password,email blank
// @Title UpdateCompany
// @Description Update company Info
// @Success 200 {string}
// @Failure 403
// @router /updateCompany [post]
func (u *CompanyController) UpdateCompany() {
	comapnyId := u.GetString("comapnyID")
	comapnyid,_ := strconv.Atoi(comapnyId)
	var company models.Company = models.Company{
		Id: comapnyid,
		Name: u.GetString("name"),
		Password: u.GetString("password"),
		EMail: u.GetString("email")}
	company.Update()
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