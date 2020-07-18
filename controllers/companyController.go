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
func (c *CompanyController) Signup() {
	var company models.Company = models.Company{
		Id: 0,
		Name: c.GetString("name"),
		Password: c.GetString("password"),
		EMail: c.GetString("email")}
	company.Create()
}

// tested and successfully delete a company in the database
// @Title DeleteCompany
// @Description Delete company by company ID
// @Success 200 {string}
// @Failure 403
// @router /deleteCompany [post]
func (c *CompanyController) DeleteCompany() {
	companyId := c.GetString("companyID")
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
func (c *CompanyController) UpdateCompany() {
	comapnyId := c.GetString("comapnyID")
	comapnyid,_ := strconv.Atoi(comapnyId)
	var company models.Company = models.Company{
		Id: comapnyid,
		Name: c.GetString("name"),
		Password: c.GetString("password"),
		EMail: c.GetString("email")}
	company.Update()
}

// tested and successfully created a user in the database
// the id also atomatically increased
// @Title CompanyLikesUser
// @Description Company likes a user
// @Success 200 {string}
// @Failure 403
// @router /companyLikesUser [post]
func (c *CompanyController) CompanyLikesUser() {
	comapnyId := c.GetString("comapnyID")
	comapnyid,_ := strconv.Atoi(comapnyId)
	var company = models.Company{Id:comapnyid}
	userId := c.GetString("userID")
	userid,_ := strconv.Atoi(userId)
	var user = models.User{Id:userid}
	var companyFavoriteUser models.CompanyFavoriteUser = models.CompanyFavoriteUser{
		Id: 0,
		Company: &company,
		User: &user}
	companyFavoriteUser.Create()
}

// tested and successfully created a user in the database ?????
// TODO We need to find the object from the given 2 id's and then delete the object of the right id
// the id also atomatically increased
// @Title CompanyStopLikesUser
// @Description Company stop liking a user
// @Success 200 {string}
// @Failure 403
// @router /companyStopLikesUser [post]
func (c *CompanyController) CompanyStopLikesUser() {
	comapnyId := c.GetString("comapnyID")
	comapnyid,_ := strconv.Atoi(comapnyId)
	var company = models.Company{Id:comapnyid}
	userId := c.GetString("userID")
	userid,_ := strconv.Atoi(userId)
	var user = models.User{Id:userid}
	var companyFavoriteUser models.CompanyFavoriteUser = models.CompanyFavoriteUser{
		Company: &company,
		User: &user}
	println("the id is:    ")
	println(companyFavoriteUser.Id)
	println(companyFavoriteUser.User.Id)
	println(companyFavoriteUser.Company.Id)
	companyFavoriteUser.Delete()
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