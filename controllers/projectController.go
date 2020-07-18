package controllers

import "github.com/astaxie/beego"

type ProjectController struct {
	beego.Controller
}

//tested and successfully created a user in the database
//the id also aromatically increased
// @Title Create
// @Description
// @Success 200 {string}
// @Failure 403
// @router /create [post]
func (p *ProjectController) Create() {

}