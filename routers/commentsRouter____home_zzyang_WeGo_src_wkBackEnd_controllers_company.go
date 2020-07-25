package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

    beego.GlobalControllerRouter["wkBackEnd/controllers/company:CompanyController"] = append(beego.GlobalControllerRouter["wkBackEnd/controllers/company:CompanyController"],
        beego.ControllerComments{
            Method: "CompanyLikesUser",
            Router: `/companyLikesUser`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["wkBackEnd/controllers/company:CompanyController"] = append(beego.GlobalControllerRouter["wkBackEnd/controllers/company:CompanyController"],
        beego.ControllerComments{
            Method: "CompanyStopLikesUser",
            Router: `/companyStopLikesUser`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["wkBackEnd/controllers/company:CompanyController"] = append(beego.GlobalControllerRouter["wkBackEnd/controllers/company:CompanyController"],
        beego.ControllerComments{
            Method: "DeleteCompany",
            Router: `/deleteCompany`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["wkBackEnd/controllers/company:CompanyController"] = append(beego.GlobalControllerRouter["wkBackEnd/controllers/company:CompanyController"],
        beego.ControllerComments{
            Method: "Login",
            Router: `/login`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["wkBackEnd/controllers/company:CompanyController"] = append(beego.GlobalControllerRouter["wkBackEnd/controllers/company:CompanyController"],
        beego.ControllerComments{
            Method: "Signup",
            Router: `/signup`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["wkBackEnd/controllers/company:CompanyController"] = append(beego.GlobalControllerRouter["wkBackEnd/controllers/company:CompanyController"],
        beego.ControllerComments{
            Method: "UpdateCompany",
            Router: `/updateCompany`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

}
