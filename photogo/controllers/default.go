package controllers

import (
	"github.com/astaxie/beego"
	"path"
	"fmt"
	"github.com/yutian1264/sky/com/utils/upload"
	"github.com/gin-gonic/gin/json"
	"math/rand"
	"photogo/models"
	"github.com/yutian1264/sky/com/utils/dbUtils"

	"strconv"
	//"strings"
	"strings"
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
func (c *MainController)AddLocalImgs(){
	path:=c.GetString("path")
	//fmt.Println(path)
	suffix := []string{".PNG", ".JPG"}
	files,_:=models.OpenFilesAndGetMsg(path,suffix)
	if len(files)>0{
		old:="\\"
		newStr:="/"
		for _,d:=range files{
			lat:=strconv.FormatFloat(d.Lat, 'E', -1, 64)//float64
			lng:=strconv.FormatFloat(d.Lng, 'E', -1, 64)//float64
			name:=strings.Replace(d.Name,old,newStr,-1)
			path:=name
			sql:="insert into resource (name,path,type,resource_time,createTime,remark,lat,lng,position) values" +
				"('"+name+"','"+path+"','','"+d.CreateTime+"','"+models.GetCurrentTime()+"','',"+lat+","+lng+",'"+d.City+"')"
				utils.Add(sql)
			}
	}
	c.Ctx.WriteString(path)
}
// @router /getListByPage [get]
func (this *MainController)GetListByPage(){
	res_type:=this.GetString("type")
	pageSize,_:=strconv.Atoi(this.GetString("pageSize"))
	pageCount,_:=strconv.Atoi(this.GetString("pageCount"))
	condition:="type='"+res_type+"'"
	m:=utils.QueryByPage("resource",condition,pageSize,pageCount)
	b,_:=json.Marshal(m)
	this.Ctx.WriteString(string(b))

}
// @router /getAllList [get]
func (this *MainController) GetAllList(){
	res_type:=this.GetString("type")
	m,_:=utils.Query("select* from resource where type='"+res_type+"'")
	b,_:=json.Marshal(m)
	this.Ctx.WriteString(string(b))
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