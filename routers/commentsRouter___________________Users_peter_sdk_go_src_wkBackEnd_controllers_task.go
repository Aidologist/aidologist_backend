package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

    beego.GlobalControllerRouter["wkBackEnd/controllers/task:TaskController"] = append(beego.GlobalControllerRouter["wkBackEnd/controllers/task:TaskController"],
        beego.ControllerComments{
            Method: "GetTask",
            Router: `/get`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["wkBackEnd/controllers/task:TaskController"] = append(beego.GlobalControllerRouter["wkBackEnd/controllers/task:TaskController"],
        beego.ControllerComments{
            Method: "PublishTask",
            Router: `/publish`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["wkBackEnd/controllers/task:TaskController"] = append(beego.GlobalControllerRouter["wkBackEnd/controllers/task:TaskController"],
        beego.ControllerComments{
            Method: "Update",
            Router: `/update`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

}
