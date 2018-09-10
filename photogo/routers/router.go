package routers

import (
	"photogo/controllers"
	"github.com/astaxie/beego"
	"fmt"
	"github.com/astaxie/beego/context"
)

func init() {
    beego.Router("/", &controllers.MainController{})
	ns := beego.NewNamespace("/f",
		//此处正式版时改为验证加密请求
		//beego.NSBefore(func(ctx *context.Context) {
		//	fmt.Println("这是前置过滤函数")
		//	fmt.Println(ctx.Input.Session("sid"))
		//}),
		//在此验证
		beego.NSCond(func(ctx *context.Context) bool {
			//if ua := ctx.Input.RequestBody.UserAgent(); ua != "" {
			//   return true
			//}
			fmt.Println("验证")
			return true
		}),
		beego.NSNamespace("/d",
			beego.NSInclude(
				&controllers.MainController{},
				&controllers.DirectoryManager{},
			),
		),
		//beego.NSNamespace("apk",beego.NSInclude(nil)),
	)
	beego.AddNamespace(ns)
}
