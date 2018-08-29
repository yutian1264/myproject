package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

	beego.GlobalControllerRouter["photogo/controllers:MainController"] = append(beego.GlobalControllerRouter["photogo/controllers:MainController"],
		beego.ControllerComments{
			Method: "AddLocalImgs",
			Router: `/addLocalImgs`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["photogo/controllers:MainController"] = append(beego.GlobalControllerRouter["photogo/controllers:MainController"],
		beego.ControllerComments{
			Method: "Upload",
			Router: `/addphoto`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["photogo/controllers:MainController"] = append(beego.GlobalControllerRouter["photogo/controllers:MainController"],
		beego.ControllerComments{
			Method: "GetAllList",
			Router: `/getAllList`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["photogo/controllers:MainController"] = append(beego.GlobalControllerRouter["photogo/controllers:MainController"],
		beego.ControllerComments{
			Method: "Tpost",
			Router: `/testpost`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

}
