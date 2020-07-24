package controllers

import (
	"github.com/astaxie/beego"
	"strconv"
	"time"
	"wkBackEnd/models"
)

// Operations about Tasks
type TaskController struct {
	beego.Controller
}

// This is tested successfully
// You can use chinese as the part of the string but you can't ues emoji
// @Title Post
// @Description Company posting a task
// @Success 200 {string}
// @Failure 403
// @router /post [post]
func (p *TaskController) Post() {
	var company = new(models.Company)
	id, _ := strconv.Atoi(p.GetString("id"))
	company.Read(id)
	var layout string = "2006-01-02 15:04:05"
	EndTime, _ := time.Parse(layout, p.GetString("endtime"))
	DueTime, _ := time.Parse(layout, p.GetString("duetime"))
	Payment, _ := strconv.Atoi(p.GetString("payment"))
	var task = models.Task {
		Title: p.GetString("title"),
		Desc: p.GetString("desc"),
		Payment:Payment,
		PayCurrency: p.GetString("payCurrency"),
		PublishTime: time.Now(),
		UpdateTime: time.Now(),
		EndTime: EndTime,
		DueTime: DueTime,
		CompanyWhoPosted: company}
	company.TasksPosted = append(company.TasksPosted, &task)
	task.Create()
	company.Update()
}

// This is successfully tested
// @Title GetAllPostedTask
// @Description Get all tasks posted by the company
// @Success 200 {string}
// @Failure 403
// @router /get/allposted [post]
func (p *TaskController) GetAllPostedTask() {
	var company = new(models.Company)
	id, _ := strconv.Atoi(p.GetString("id"))
	company.Read(id)
	company.ReadPostedTaskList()
	var companyTaskMap = make(map[int]models.Task)
	//for i := range company.Projects {
	//	println(company.Projects[i])
	//}
	for i := range company.TasksPosted {
		companyTaskMap[i+1] = *company.TasksPosted[i]
	}
	p.Data["json"] = map[string]interface{}{
		"size" : len(company.TasksPosted),
		"tasks" : companyTaskMap}
	p.ServeJSON()
}

// This is tested successfully
// You can not use post as the function type
// You have to use delete as the type
// @Title Delete
// @Description
// @Success 200 {string}
// @Failure 403
// @router /delete [delete]
func (p *TaskController) Delete() {
	var task = new(models.Task)
	id, _ := strconv.Atoi(p.GetString("id"))
	task.Read(id)
	task.Delete()
}

// Tested successfully
// Todo be careful when you send the jason request
// TODO all the fields have to be filled for the database to update the object
// TODO if you only want to update one field, just send the request and leave all other fields the same
// @Title Update
// @Description
// @Success 200 {string}
// @Failure 403
// @router /update [post]
func (t *TaskController) Update() {
	pay, _ := strconv.Atoi(t.GetString("payment"))
	id, _ := strconv.Atoi(t.GetString("id"))
	var layout string = "2006-01-02 15:04:05"
	endtime, _ := time.Parse(layout, t.GetString("endtime"))
	duetime, _ := time.Parse(layout, t.GetString("duetime"))
	var task models.Task
	task.Read(id)
	task.Title = t.GetString("title")
	task.Desc = t.GetString("desc")
	task.Payment = pay
	task.PayCurrency = t.GetString("paycurrency")
	task.UpdateTime = time.Now()
	task.EndTime = endtime
	task.DueTime = duetime
	task.Update()
}

// @Title GetTask
// @Description
// @Success 200 {string}
// @Failure 403
// @router /get [get]
func (t *TaskController) GetTask() {
	var task models.Task
	task.ReadTask()
	t.Data["json"] = &task
	t.ServeJSON()
}