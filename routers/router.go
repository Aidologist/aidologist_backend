// @APIVersion 1.0.0
// @Title beego Test API
// @Description beego has a very cool tools to autogenerate documents for your API
// @Contact astaxie@gmail.com
// @TermsOfServiceUrl http://beego.me/
// @License Apache 2.0
// @LicenseUrl http://www.apache.org/licenses/LICENSE-2.0.html
package routers

import (
	"github.com/astaxie/beego"
	companyController "wkBackEnd/controllers/company"
	projectController "wkBackEnd/controllers/company/project"
	taskController "wkBackEnd/controllers/task"
	userController "wkBackEnd/controllers/user"
)

func init() {
	ns := beego.NewNamespace("/main",
		beego.NSNamespace("/user",
			beego.NSInclude(
				&userController.UserController{},
			),
		),
		beego.NSNamespace("/task",
			beego.NSInclude(
				&taskController.TaskController{},
			),
		),
		beego.NSNamespace("/company",
			beego.NSInclude(
				&companyController.CompanyController{},
			),
			beego.NSNamespace("/project",
				beego.NSInclude(
					&projectController.ProjectController{},
				),
			),
		),
	)
	beego.AddNamespace(ns)
}
