package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

    beego.GlobalControllerRouter["wkBackEnd/controllers/user:UserController"] = append(beego.GlobalControllerRouter["wkBackEnd/controllers/user:UserController"],
        beego.ControllerComments{
            Method: "DeleteUser",
            Router: `/deleteUser`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["wkBackEnd/controllers/user:UserController"] = append(beego.GlobalControllerRouter["wkBackEnd/controllers/user:UserController"],
        beego.ControllerComments{
            Method: "Login",
            Router: `/login`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["wkBackEnd/controllers/user:UserController"] = append(beego.GlobalControllerRouter["wkBackEnd/controllers/user:UserController"],
        beego.ControllerComments{
            Method: "Signup",
            Router: `/signup`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["wkBackEnd/controllers/user:UserController"] = append(beego.GlobalControllerRouter["wkBackEnd/controllers/user:UserController"],
        beego.ControllerComments{
            Method: "UpdateUser",
            Router: `/updateUser`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

}
