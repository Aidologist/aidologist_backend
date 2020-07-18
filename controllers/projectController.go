package controllers

import (
	"github.com/astaxie/beego"
	"strconv"
	"time"
	"wkBackEnd/models"
)

type ProjectController struct {
	beego.Controller
}

// @Title Create
// @Description
// @Success 200 {string}
// @Failure 403
// @router /create [post]
func (p *ProjectController) Create() {
	var company = new(models.Company)
	id, _ := strconv.Atoi(p.GetString("id"))
	company.Read(id)
	var layout string = "2006-01-02 15:04:05"
	endtime, _ := time.Parse(layout, p.GetString("endtime"))
	duetime, _ := time.Parse(layout, p.GetString("duetime"))
	var project = models.Project {
		Title: p.GetString("title"),
		Desc: p.GetString("desc"),
		Cover: p.GetString("cover"),
		EndTime: endtime,
		DueTime: duetime,
		Company: company}
	company.Projects = append(company.Projects, &project)
	project.Create()
	company.Update()
}

// @Title Create
// @Description
// @Success 200 {string}
// @Failure 403
// @router /get/all [post]
func (p *ProjectController) GetAll() {
	var company = new(models.Company)
	id, _ := strconv.Atoi(p.GetString("id"))
	company.Read(id)
	//println(company.Id)
	//println(company.EMail)
	//println(company.Name)
	//println(company.Password)
	//println(company.Projects)
	company.ReadAllProject()
	var companyProjectsMap = make(map[int]models.Project)
	//for i := range company.Projects {
	//	println(company.Projects[i])
	//}
	for i := range company.Projects {
		companyProjectsMap[i+1] = *company.Projects[i]
	}
	p.Data["json"] = map[string]interface{}{
		"size" : len(company.Projects),
		"projects" : companyProjectsMap}
	p.ServeJSON()
}