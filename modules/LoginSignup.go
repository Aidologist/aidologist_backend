package modules

import (
	"github.com/astaxie/beego"
	"wkBackEnd/models"
	"wkBackEnd/mooPacks/moohttp"
)

func Login(username, password, email string, c beego.Controller)  {
	var json map[string]bool = make(map[string]bool)
	json["success"] = false
	var http = moohttp.MooRequest{CtxInput: c.Ctx.Input.RequestBody}
	http.RequestBodyToMap()
	var user models.User = models.User{
		Id: 0,
		Password: http.ReadString("password"),
		EMail: http.ReadString("email"),
	}
	user.ReadByEmail()
	c.SetSession("id", user.Id)
	c.SetSession("username", user.Username)
	c.SetSession("password", password)
	c.SetSession("email", email)
	json["success"] = true
	c.Data["json"] = json
	c.ServeJSON()
}

func Signup(username, password, email string, c beego.Controller) {
	var json map[string]bool = make(map[string]bool)
	json["success"] = false
	var http = moohttp.MooRequest{CtxInput: c.Ctx.Input.RequestBody}
	http.RequestBodyToMap()
	var user models.User = models.User{
		Id: 0,
		Username: http.ReadString("username"),
		Password: http.ReadString("password"),
		EMail: http.ReadString("email"),
	}
	user.Create()
	c.SetSession("id", user.Id)
	c.SetSession("username", username)
	c.SetSession("password", password)
	c.SetSession("email", email)
	json["success"] = true
	c.Data["json"] = json
	c.ServeJSON()
}

func VerifyEmail(email string, c beego.Controller)  {
	var json map[string]bool = make(map[string]bool)
	json["emailRegistered"] = false
	var http = moohttp.MooRequest{CtxInput: c.Ctx.Input.RequestBody}
	http.RequestBodyToMap()
	var user models.User = models.User{
		Id: 0,
		EMail: http.ReadString("email"),
	}
	if user.CheckIfEmailExists() {
		json["emailRegistered"] = true
	}
	c.Data["json"] = json
	c.ServeJSON()
}

// TODO: verify is the code [ c.GetSession("id").(int) ] works
func VerifyLogin(c beego.Controller)  {
	var json map[string]interface{} = make(map[string]interface{})
	json["isLogin"] = false
	if c.GetSession("id") != nil {
		json["isLogin"] = true
		json["userId"] = c.GetSession("id").(int)
		json["userInfo"] = map[string]string{
			"username" : c.GetSession("username").(string),
			"password" : c.GetSession("password").(string),
			"email" : c.GetSession("email").(string),
		}
	}
	c.Data["json"] = json
	c.ServeJSON()
}

func DeleteSession(c beego.Controller) {
	c.DelSession("id")
	c.DelSession("username")
	c.DelSession("password")
	c.DelSession("email")
}