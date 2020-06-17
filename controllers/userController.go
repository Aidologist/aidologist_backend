package controllers

import (
	"github.com/astaxie/beego"
	"wkBackEnd/models"
)

// Operations about Users
type UserController struct {
	beego.Controller
}

// @Title Signup
// @Description
// @Success 200 {string}
// @Failure 403
// @router /signup [post]
func (u *UserController) Signup() {
	var user models.User = models.User{
		Id: 0,
		Username: u.GetString("username"),
		Password: u.GetString("password"),
		EMail: u.GetString("email")}
	user.AddUser()
}

// @Title Login
// @Description
// @Success 200 {string}
// @Failure 403
// @router /login [post]
func (u *UserController) Login() {
	var user *models.User = new(models.User)
	user.Username = u.GetString("username")
	user.EMail = u.GetString("email")
	user.GetUser()
	if user.Password == u.GetString("password") {
		u.Data["info"] = "success"
	} else {
		u.Data["info"] = "fail"
	}
}