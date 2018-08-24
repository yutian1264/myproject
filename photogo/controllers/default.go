package controllers

import (
	"github.com/astaxie/beego"
	"path"
	"fmt"
	"sky/com/utils/upload"
	"github.com/gin-gonic/gin/json"
)

type MainController struct {
	beego.Controller
}

func (c *MainController) Get() {
	c.Data["Website"] = "beego.me"
	c.Data["Email"] = "astaxie@gmail.com"
	c.TplName = "upload.html"
}
// @router /addphoto [post]
func (this *MainController) Upload() {
	this.TplName="upload.html"
	//go uploadThread(this)
	var ch=make(chan int)
	go upload.UploadBreakPoint(this.Ctx.Request,"file","static/upload/",ch)
	c:=<-ch
	var result=make(map[string]interface{})
	result["code"]=200
	result["status"]=c
	b,_:=json.Marshal(result)
	this.Ctx.WriteString(string(b))
}
func uploadThread(this *MainController){
	f, h, _ := this.GetFile("file")	//获取上传的文件

	f.Close()
	//	this.SaveToFile("image", path+ h.Filename)	//存文件    WaterMark(path)	//给文件加水印
	error:=this.SaveToFile("image", path.Join("/static/uploadfile",h.Filename))
	fmt.Println(error,h.Filename)
	this.Ctx.WriteString("OK")
	this.TplName="upload.html"
}