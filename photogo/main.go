package main

import (
	_ "photogo/routers"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/session"
	"fmt"
	"github.com/yutian1264/sky/com/utils/dbUtils"
	"photogo/models"
	"runtime"
)
var globalSessions *session.Manager
func init() {
	//err:=utils.InitMySQLDB("root","root","localhost:3306","farming")
	err:=utils.InitMySQLDB("root","wang","192.168.248.138:3306","tfarming")
	if err!=nil{
		fmt.Println("mysql connected error")
		return
	}
	//utils.MongodbInit("ec","192.168.180.115","27017","ec","panda","ec",4096)
	sessionConfig := &session.ManagerConfig{
		CookieName:"gosessionid",
		EnableSetCookie: true,
		Gclifetime:3600,
		Maxlifetime: 3600,
		Secure: false,
		CookieLifeTime: 3600,
		ProviderConfig: "./tmp",
	}

	globalSessions, _ = session.NewManager("memory",sessionConfig)
	go globalSessions.GC()


}
func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	beego.BConfig.WebConfig.StaticDir["/static"] = "static"
	beego.BConfig.WebConfig.StaticDir["/views"] = "views"
	beego.SetStaticPath( "/j","J:/")
	beego.SetStaticPath( "/h","H:/")
	beego.SetStaticPath( "/d","D:/")
	beego.SetStaticPath( "/c","C:/")
	beego.SetStaticPath( "/e","E:/")
	beego.SetStaticPath( "/g","G:/")
	beego.BConfig.WebConfig.Session.SessionOn = true
	beego.InsertFilter("*", beego.BeforeRouter, models.Allow(&models.Options{
		AllowAllOrigins:  true,
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Authorization", "Access-Control-Allow-Origin", "Access-Control-Allow-Headers", "Content-Type"},
		ExposeHeaders:    []string{"Content-Length", "Access-Control-Allow-Origin", "Access-Control-Allow-Headers", "Content-Type"},
		AllowCredentials: true,
	}))
	//初始化sql 配置文件
	utils.Init("/static/cont/datasource.xml")
	beego.Run()
}


