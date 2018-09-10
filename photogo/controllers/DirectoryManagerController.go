package controllers

import (
	"github.com/astaxie/beego"
	"github.com/yutian1264/sky/com/utils/dbUtils"
	"encoding/json"
	"photogo/models"
)

type DirectoryManager struct {
	beego.Controller
}
// @router /getDirectoryListByUser [get]
func (this *DirectoryManager)GetDirectoryListByUser(){
	user:=this.GetString("user")
	m,_:=utils.Query("select *from directory where userName='"+user+"'")
	b,_:=json.Marshal(m)
	this.Ctx.WriteString(string(b))
}
// @router /addDirectory [post]
func (this *DirectoryManager)AddDirectory(){
	param:=this.Ctx.Request.Form.Get("param")
	m:=make(map[string]interface{})
	json.Unmarshal([]byte(param),&m)
	sql:="insert into directory (name,userName,createTime) values('"+m["name"].(string)+"','"+m["user"].(string)+"','"+models.GetCurrentTime()+"')"
	res:=utils.Add(sql)
	//b,_:=json.Marshal(m)
	this.Ctx.WriteString(res)
}
