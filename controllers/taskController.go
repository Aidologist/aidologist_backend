package controllers

import (
	"github.com/astaxie/beego"
	"strconv"
	"time"
	"wkBackEnd/models"
)

// Operations about Users
type TaskController struct {
	beego.Controller
}

// @Title PublishTask
// @Description
// @Success 200 {string}
// @Failure 403
// @router /publish [post]
func (u *TaskController) PublishTask() {
	pay, _ := strconv.Atoi(u.GetString("payment"))
	var layout string = "2006-01-02 15:04:05"
	endtime, _ := time.Parse(layout, u.GetString("endtime"))
	duetime, _ := time.Parse(layout, u.GetString("duetime"))
	var task models.Task = models.Task{
		Id: 0,
		Title: u.GetString("title"),
		Desc: u.GetString("desc"),
		Payment: pay,
		PayCurrency: u.GetString("paycurrency"),
		PublishTIme: time.Now(),
		UpdateTime: time.Now(),
		EndTime: endtime,
		DueTime: duetime}
	task.AddTask()
}

// @Title Update
// @Description
// @Success 200 {string}
// @Failure 403
// @router /update [post]
func (u *TaskController) Update() {
	pay, _ := strconv.Atoi(u.GetString("payment"))
	id, _ := strconv.Atoi(u.GetString("id"))
	var layout string = "2006-01-02 15:04:05"
	endtime, _ := time.Parse(layout, u.GetString("endtime"))
	duetime, _ := time.Parse(layout, u.GetString("duetime"))
	var task models.Task
	task.ReadTaskbyId(id)
	task.Title = u.GetString("title")
	task.Desc = u.GetString("desc")
	task.Payment = pay
	task.PayCurrency = u.GetString("paycurrency")
	task.UpdateTime = time.Now()
	task.EndTime = endtime
	task.DueTime = duetime
	task.UpdateTask()
}

// @Title GetTask
// @Description
// @Success 200 {string}
// @Failure 403
// @router /get [get]
func (u *TaskController) GetTask() {
	var task models.Task
	task.ReadTask()
	u.Data["json"] = &task
	u.ServeJSON()
}