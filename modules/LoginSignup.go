package modules

import (
	"github.com/astaxie/beego"
)

func Login()  {

}

func Signup(username, password, email string, c beego.Controller) {
	var json map[string]bool = make(map[string]bool)
	c.SetSession("username", username)
	c.SetSession("password", password)
	c.SetSession("email", email)
	c.Data["json"] = json
	c.ServeJSON()
}

func VerifyEmail(email string, c beego.Controller)  {

}

func VerifyLogin()  {

}