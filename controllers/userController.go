package controllers

import (
	"github.com/astaxie/beego"
	"wkBackEnd/modules"
	"wkBackEnd/mooPacks/moohttp"
)

// Operations about Users
type UserController struct {
	beego.Controller
}

func (c *UserController) Signup() {
	var http = moohttp.MooRequest{CtxInput: c.Ctx.Input.RequestBody}
	http.RequestBodyToMap()
	modules.Signup(
		http.ReadString("username"),
		http.ReadString("password"),
		http.ReadString("email"),
		c.Controller,
	)
}