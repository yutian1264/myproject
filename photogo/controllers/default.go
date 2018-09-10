package controllers

import (
	"github.com/astaxie/beego"
	"path"
	"fmt"
	"github.com/yutian1264/sky/com/utils/upload"
	"encoding/json"
	"photogo/models"
	dbUtils "github.com/yutian1264/sky/com/utils/dbUtils"
	. "github.com/yutian1264/sky/com/utils"
	"strconv"
	//"strings"
	"strings"
	"log"
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
	go getImgs(c)
}
//var group sync.WaitGroup
var pool models.Pool
func getImgs(c *MainController){

	path:=c.GetString("path")
	//fmt.Println(path)
	suffix := []string{".PNG", ".JPG"}
	files,_:=models.OpenFilesAndGetMsg(path,suffix)
	if len(files)>0{
		pool.Init(7, len(files));
		old:="\\"
		newStr:="/"
		for _,d:=range files{
			//group.Add(1)
			lat:=strconv.FormatFloat(d.Lat, 'E', -1, 64)//float64
			lng:=strconv.FormatFloat(d.Lng, 'E', -1, 64)//float64
			path:=strings.Replace(d.Name,old,newStr,-1)
			name:=path
			name=name[strings.LastIndex(name,"/")+1:]
			temp:=strings.Split(path,"/")
			thumb:="thumb_"+name
			thumbSaveParh:="H:/thumb/"+temp[len(temp)-2]
			b,err:=PathNotExistsCreate(thumbSaveParh)
			if !b{
				recover()
				log.Println("mkdir failed![%v]\n", err)
			}
			sql:="insert into resource (name,path,type,resource_time,createTime,remark,lat,lng,position,thumb) values" +
				"('"+name+"','"+path+"','','"+d.CreateTime+"','"+models.GetCurrentTime()+"','',"+lat+","+lng+",'"+d.City+"','"+thumbSaveParh+"/"+thumb+"')"
			pool.AddTask(func()error {
				return createThumb(path,thumbSaveParh,thumb,sql);
			});

		}
		pool.SetFinishCallback(func() {
			fmt.Println("更新完成")
			c.Ctx.WriteString("OK")
		});
		pool.Start();
		//group.Wait()
		pool.Stop();
	}else{
		c.Ctx.WriteString("")
	}

}

func createThumb(path,thumbSaveParh,thumbName,sql string)error{
	dbUtils.Add(sql)
	err:=CreateThumb(path,thumbSaveParh,thumbName,500,500)
	if err!=nil{
		fmt.Println(err)
	}
	//group.Done()
	return err
}
// @router /UpdateRess [post]
func (this *MainController)UpdateRess(){
	this.Ctx.Request.ParseForm()
	param:=this.Ctx.Request.Form.Get("param")
	m:=make(map[string]interface{})
	json.Unmarshal([]byte(param),&m)
	owner:=m["owner"].(string)
	res_type:=m["type"].(string)
	ids:=m["ids"].([]interface{})
	fmt.Println(owner,res_type,ids)
	for _,s:=range ids{
		fmt.Println(s)
		sql:="update resource set sub_user='"+owner +"',type='"+res_type+"' where id="+s.(string)
		dbUtils.Sql_update(sql,false)
	}
	fmt.Println(m)
	this.Ctx.WriteString("")

}
// @router /getListByPage [get]
func (this *MainController)GetListByPage(){
	res_type:=this.GetString("type")
	pageSize,_:=strconv.Atoi(this.GetString("pageSize"))
	pageCount,_:=strconv.Atoi(this.GetString("pageCount"))
	var condition="";
	if len(res_type)==0{
		condition="sub_user is null"
	}else{
		condition="sub_user='"+res_type+"'"
	}

	m:=dbUtils.QueryByPage("resource",condition,pageSize,pageCount)
	b,_:=json.Marshal(m)
	this.Ctx.WriteString(string(b))

}
// @router /getAllList [get]
func (this *MainController) GetAllList(){
	res_type:=this.GetString("type")
	m,_:=dbUtils.Query("select* from resource where type='"+res_type+"'")
	b,_:=json.Marshal(m)
	this.Ctx.WriteString(string(b))
}
// @router /getResTypeList [get]
func(this *MainController)GetResTypeList(){
	m,_:=dbUtils.Query("select * from res_type")
	b,_:=json.Marshal(m)
	this.Ctx.WriteString(string(b))
}
// @router /getResOwnerList [get]
func(this *MainController)GetResOwnerList(){
	m,_:=dbUtils.Query("select * from res_owner where level is null")
	b,_:=json.Marshal(m)
	this.Ctx.WriteString(string(b))
}
// @router /addphoto [post]
func (this *MainController) Upload() {
	this.TplName="upload.html"
	//go uploadThread(this)
	var ch=make(chan upload.ResultChan)
	this.Ctx.Request.ParseForm()
	savePath:=this.Ctx.Request.Form.Get("savePath")
	go upload.UploadBreakPoint(this.Ctx.Request,"file",savePath,ch)
	c:=<-ch
	var result=make(map[string]interface{})
	result["code"]=200
	result["status"]=c.Status
	result["data"]=c
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
// @router /addResource [post]
func (this *MainController)AddResource(){
	this.Ctx.Request.ParseForm()
	param:=this.Ctx.Request.Form.Get("param")
	m:=make(map[string]interface{})
	json.Unmarshal([]byte(param),&m)
	msg:=m["fileMsg"].([]interface{})
	fmt.Println(msg)
	res_type:=m["type"].(string)
	lat:=strconv.FormatFloat(m["lat"].(float64), 'E', -1, 64)//float64
	lng:=strconv.FormatFloat(m["lng"].(float64), 'E', -1, 64)//float64
	position:=m["position"].(string)
	user:=m["user"].(string)
	category:=m["category"].(string)
	for index,item:=range msg{
		tmp:=item.(map[string]interface{})["data"].(map[string]interface{})
		fmt.Println(index,tmp["FileName"])
		sql:="insert into resource (name,path,type,createTime,sub_user,lat,lng,thumb,position,category)" +
			"values('"+tmp["FileName"].(string)+"','"+tmp["FilePath"].(string)+"','"+res_type+"'," +
				"'"+models.GetCurrentTime()+"','"+user+"','"+lat+"','"+lng+"','"+tmp["ThumbPath"].(string)+"','"+position+"','"+category+"')"
		fmt.Println(sql)
		s:=dbUtils.Add(sql)
		fmt.Println("result:%s",s)
	}
	this.Ctx.WriteString("OK")
}