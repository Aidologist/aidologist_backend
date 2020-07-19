// @APIVersion 1.0.0
// @Title beego Test API
// @Description beego has a very cool tools to autogenerate documents for your API
// @Contact astaxie@gmail.com
// @TermsOfServiceUrl http://beego.me/
// @License Apache 2.0
// @LicenseUrl http://www.apache.org/licenses/LICENSE-2.0.html
package routers

import (
	"wkBackEnd/controllers/user"

	"github.com/astaxie/beego"
)

func init() {
	ns := beego.NewNamespace("/main",
		beego.NSNamespace("/user",
			beego.NSInclude(
				&user.UserController{},
			),
		),
		beego.NSNamespace("/task",
			beego.NSInclude(
				&task.TaskController{},
			),
		),
		beego.NSNamespace("/company",
			beego.NSInclude(
				&company.CompanyController{},
			),
			beego.NSNamespace("/project",
				beego.NSInclude(
					&project.ProjectController{},
				),
			),
		),
	)
	beego.AddNamespace(ns)
}
