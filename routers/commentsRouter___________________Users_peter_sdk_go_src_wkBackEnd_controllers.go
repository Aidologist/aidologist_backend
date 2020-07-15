package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

    beego.GlobalControllerRouter["wkBackEnd/controllers:CompanyController"] = append(beego.GlobalControllerRouter["wkBackEnd/controllers:CompanyController"],
        beego.ControllerComments{
            Method: "DeleteCompany",
            Router: `/deleteCompany`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["wkBackEnd/controllers:CompanyController"] = append(beego.GlobalControllerRouter["wkBackEnd/controllers:CompanyController"],
        beego.ControllerComments{
            Method: "Login",
            Router: `/login`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["wkBackEnd/controllers:CompanyController"] = append(beego.GlobalControllerRouter["wkBackEnd/controllers:CompanyController"],
        beego.ControllerComments{
            Method: "Signup",
            Router: `/signup`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["wkBackEnd/controllers:CompanyController"] = append(beego.GlobalControllerRouter["wkBackEnd/controllers:CompanyController"],
        beego.ControllerComments{
            Method: "UpdateCompany",
            Router: `/updateCompany`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["wkBackEnd/controllers:TaskController"] = append(beego.GlobalControllerRouter["wkBackEnd/controllers:TaskController"],
        beego.ControllerComments{
            Method: "GetTask",
            Router: `/get`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["wkBackEnd/controllers:TaskController"] = append(beego.GlobalControllerRouter["wkBackEnd/controllers:TaskController"],
        beego.ControllerComments{
            Method: "PublishTask",
            Router: `/publish`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["wkBackEnd/controllers:TaskController"] = append(beego.GlobalControllerRouter["wkBackEnd/controllers:TaskController"],
        beego.ControllerComments{
            Method: "Update",
            Router: `/update`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["wkBackEnd/controllers:UserController"] = append(beego.GlobalControllerRouter["wkBackEnd/controllers:UserController"],
        beego.ControllerComments{
            Method: "DeleteUser",
            Router: `/deleteUser`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["wkBackEnd/controllers:UserController"] = append(beego.GlobalControllerRouter["wkBackEnd/controllers:UserController"],
        beego.ControllerComments{
            Method: "Login",
            Router: `/login`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["wkBackEnd/controllers:UserController"] = append(beego.GlobalControllerRouter["wkBackEnd/controllers:UserController"],
        beego.ControllerComments{
            Method: "Signup",
            Router: `/signup`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["wkBackEnd/controllers:UserController"] = append(beego.GlobalControllerRouter["wkBackEnd/controllers:UserController"],
        beego.ControllerComments{
            Method: "UpdateUser",
            Router: `/updateUser`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

}
