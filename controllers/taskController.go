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

// @Title PublishTask
// @Description
// @Success 200 {string}
// @Failure 403
// @router /publish [post]
func (t *TaskController) PublishTask() {
	pay, _ := strconv.Atoi(t.GetString("payment"))
	var layout string = "2006-01-02 15:04:05"
	endtime, _ := time.Parse(layout, t.GetString("endtime"))
	duetime, _ := time.Parse(layout, t.GetString("duetime"))
	var task models.Task = models.Task{
		Id:          0,
		Title:       t.GetString("title"),
		Desc:        t.GetString("desc"),
		Payment:     pay,
		PayCurrency: t.GetString("paycurrency"),
		PublishTIme: time.Now(),
		UpdateTime:  time.Now(),
		EndTime:     endtime,
		DueTime:     duetime}
	task.AddTask()
}

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
	task.ReadTaskbyId(id)
	task.Title = t.GetString("title")
	task.Desc = t.GetString("desc")
	task.Payment = pay
	task.PayCurrency = t.GetString("paycurrency")
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
func (t *TaskController) GetTask() {
	var task models.Task
	task.ReadTask()
	t.Data["json"] = &task
	t.ServeJSON()
}