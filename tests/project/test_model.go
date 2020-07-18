package testproject

import (
	"github.com/astaxie/beego/orm"
	"time"
	"wkBackEnd/models"
)

func AddTestProject01() {
	endtime, _ := time.Parse("2006-01-02 15:04:05", "2006-01-02 15:04:05")
	duetime, _ := time.Parse("2006-01-02 15:04:05", "2006-01-02 15:04:05")
	var byteDance = &models.Company{
		Name: "ByteDance",
		EMail: "ByteDance@gmail.com",
		Password: "123456789"}
	var byteDanceProjects = []*models.Project{
		&models.Project{
			Title: "TikTok",
			Desc: "TikTok is a Chinese video-sharing social networking service owned by ByteDance, " +
				"a Beijing-based internet technology company founded in 2012 by Zhang Yiming. " +
				"It is used to create short dance, lip-sync, comedy and talent videos.",
			Cover:   "https://en.wikipedia.org/wiki/TikTok#/media/File:TikTok_logo.svg",
			EndTime: endtime,
			DueTime: duetime,
			Company: byteDance},
		&models.Project{
			Title: "TouTiao",
			Desc: "Toutiao or Jinri Toutiao is a Chinese news and information content platform, " +
				"a core product of the Beijing-based company ByteDance. " +
				"By analyzing the features of content, users and users’ interaction with content, " +
				"the company's algorithm models generate a tailored feed list of content for each user.",
			Cover:   "https://en.wikipedia.org/wiki/Toutiao#/media/File:ToutiaoLogo2017.png",
			EndTime: endtime,
			DueTime: duetime,
			Company: byteDance},
	}
	byteDance.Projects = byteDanceProjects
	o := orm.NewOrm()
	o.Using("xq")
	o.Insert(byteDance)
	o.Insert(byteDanceProjects)
}