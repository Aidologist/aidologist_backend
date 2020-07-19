package controllers

import (
	"github.com/astaxie/beego"
	"strconv"
	"wkBackEnd/models"
)

// Operations about Users
type UserController struct {
	beego.Controller
}

//tested and successfully created a user in the database
//the id also atomatically increased
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
	//user.AddUser()
	user.Create()
}

// tested and successfully delete a user in the database
// @Title DeleteUser
// @Description Delete user by user ID
// @Success 200 {string}
// @Failure 403
// @router /deleteUser [post]
func (u *UserController) DeleteUser() {
	userId := u.GetString("userID")
	userid,_ := strconv.Atoi(userId)
	var user models.User = models.User{
		Id: userid}
	user.Delete()
}

// Tested sucessfully
// But you can not leave username,password,email blank
// @Title UpdateUser
// @Description Update user Info
// @Success 200 {string}
// @Failure 403
// @router /updateUser [post]
func (u *UserController) UpdateUser() {
	userId := u.GetString("userID")
	userid,_ := strconv.Atoi(userId)
	var user models.User = models.User{
		Id: userid,
		Username: u.GetString("username"),
		Password: u.GetString("password"),
		EMail: u.GetString("email")}
	user.Update()
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