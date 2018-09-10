package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

	beego.GlobalControllerRouter["photogo/controllers:DirectoryManager"] = append(beego.GlobalControllerRouter["photogo/controllers:DirectoryManager"],
		beego.ControllerComments{
			Method: "AddDirectory",
			Router: `/addDirectory`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["photogo/controllers:DirectoryManager"] = append(beego.GlobalControllerRouter["photogo/controllers:DirectoryManager"],
		beego.ControllerComments{
			Method: "GetDirectoryListByUser",
			Router: `/getDirectoryListByUser`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["photogo/controllers:MainController"] = append(beego.GlobalControllerRouter["photogo/controllers:MainController"],
		beego.ControllerComments{
			Method: "UpdateRess",
			Router: `/UpdateRess`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["photogo/controllers:MainController"] = append(beego.GlobalControllerRouter["photogo/controllers:MainController"],
		beego.ControllerComments{
			Method: "AddLocalImgs",
			Router: `/addLocalImgs`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["photogo/controllers:MainController"] = append(beego.GlobalControllerRouter["photogo/controllers:MainController"],
		beego.ControllerComments{
			Method: "AddResource",
			Router: `/addResource`,
			AllowHTTPMethods: []string{"post"},
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
			Method: "GetListByPage",
			Router: `/getListByPage`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["photogo/controllers:MainController"] = append(beego.GlobalControllerRouter["photogo/controllers:MainController"],
		beego.ControllerComments{
			Method: "GetResOwnerList",
			Router: `/getResOwnerList`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["photogo/controllers:MainController"] = append(beego.GlobalControllerRouter["photogo/controllers:MainController"],
		beego.ControllerComments{
			Method: "GetResTypeList",
			Router: `/getResTypeList`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

}
