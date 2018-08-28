package controllers

import (
	"github.com/astaxie/beego"
	"path"
	"fmt"
	"sky/com/utils/upload"
	"github.com/gin-gonic/gin/json"
	"math/rand"
)

type MainController struct {
	beego.Controller
}

func (c *MainController) Get() {
	c.Data["Website"] = "beego.me"
	c.Data["Email"] = "astaxie@gmail.com"
	c.TplName = "index.html"
}
// @router /addLocalImgs [get]
func (c *MainController) getImgMsg(){
	path:=c.GetString("path")
	fmt.Println(path)
	c.Ctx.WriteString(path)
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
// @router /testpost [post]
func (this *MainController)Tpost(){
	this.TplName="upload.html"
	this.Ctx.Request.ParseForm()
	forms:=this.Ctx.Request.Form
	fmt.Println(forms)
	i:=rand.Intn(4)
	var result="";
	switch i {
		case 0:
			result=`{
			    "code":0,
			    "message":"ok",
			    "data":
			    {
			           "action":"weather.weather_forecast",
			           "executable": 1,
			           "nlg":
			               {
			                   "text":"{地点}{时间}[天气情况数据]，气温[XX度到XX度]，[温馨提示2]or[温馨提示1]，（空气[质量]，[温馨提示3]，）[天气预警]",
			                   "type":1
			               },
			           "slot":
			            {
			                "date":"明天",
			                "location":"北京",
			                "type":"weather"
			            }
			    }
			   }`
		case 1:
			result=`{
			    "code":1,
			    "message":"error",
			    "data":
			    {
			           "action":"weather.weather_forecast",
			           "executable": 1,
			           "nlg":
			               {
			                   "text":"{地点}{时间}[天气情况数据]，气温[XX度到XX度]，[温馨提示2]or[温馨提示1]，（空气[质量]，[温馨提示3]，）[天气预警]",
			                   "type":1
			               },
			           "slot":
			            {
			                "date":"明天",
			                "location":"北京",
			                "type":"weather"
			            }
			    }
			   }`
		case 2:
			result=`{
			    "code":0,
			    "message":"terter",
			    "data":
			    {
			           "action":"weather.weather_forecast",
			           "executable": 1,
			           "slot":
			            {
			                "date":"",
			                "location":"",
			                "type":"weather"
			            }
			    }
			   }`
		case 3:
			result=`{
			    "code":0,
			    "message":"sssss",
			    "data":
			    {
			           "action":"weather.weather_forecast",
			           "executable": 1,
			           "slot":
			            {
			                "date":"",
			                "location":"北京",
			                "type":"weather"
			            }
			    }
			   }`
	default:
		result=`{
			    "code":404,
			    "message":"ok",
			    "data":
			    {
			           "action":"weather.weather_forecast",
			           "executable": 1,
			           "nlg":
			               {
			                   "text":"{地点}{时间}[天气情况数据]，气温[XX度到XX度]，[温馨提示2]or[温馨提示1]，（空气[质量]，[温馨提示3]，）[天气预警]",
			                   "type":1
			               },
			           "slot":
			            {
			                "date":"明天",
			                "location":"北京",
			                "type":"weather"
			            }
			    }
			   }`
	}
	this.Ctx.WriteString(result)
}