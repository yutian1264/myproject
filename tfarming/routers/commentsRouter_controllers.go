package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

	beego.GlobalControllerRouter["tfarming/controllers:MainController"] = append(beego.GlobalControllerRouter["tfarming/controllers:MainController"],
		beego.ControllerComments{
			Method: "Upload",
			Router: `/addphoto`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["tfarming/controllers:MainController"] = append(beego.GlobalControllerRouter["tfarming/controllers:MainController"],
		beego.ControllerComments{
			Method: "InitDb",
			Router: `/conndb`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["tfarming/controllers:MainController"] = append(beego.GlobalControllerRouter["tfarming/controllers:MainController"],
		beego.ControllerComments{
			Method: "UploadHandle",
			Router: `/upload`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["tfarming/controllers:MainController"] = append(beego.GlobalControllerRouter["tfarming/controllers:MainController"],
		beego.ControllerComments{
			Method: "UploadHandle2",
			Router: `/upload2`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

}
