package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

	beego.GlobalControllerRouter["photogo/controllers:MainController"] = append(beego.GlobalControllerRouter["photogo/controllers:MainController"],
		beego.ControllerComments{
			Method: "Upload",
			Router: `/addphoto`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

}
