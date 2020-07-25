package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

    beego.GlobalControllerRouter["wkBackEnd/controllers/company/project:ProjectController"] = append(beego.GlobalControllerRouter["wkBackEnd/controllers/company/project:ProjectController"],
        beego.ControllerComments{
            Method: "Create",
            Router: `/create`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["wkBackEnd/controllers/company/project:ProjectController"] = append(beego.GlobalControllerRouter["wkBackEnd/controllers/company/project:ProjectController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/delete`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["wkBackEnd/controllers/company/project:ProjectController"] = append(beego.GlobalControllerRouter["wkBackEnd/controllers/company/project:ProjectController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/get/all`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

}
