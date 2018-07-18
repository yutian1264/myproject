package main

import (
	_ "tfarming/routers"
	"github.com/astaxie/beego"
	dbutils "github.com/yutian1264/sky/com/utils/dbUtils"
	redisUtils  "github.com/yutian1264/sky/com/utils/redisUtils"
	utils3 "github.com/yutian1264/sky/com/utils/dbUtils"
	"os"
	"fmt"
	"encoding/json"
	"tfarming/server"
)
var redisReady bool=false
func init() {
	//初始化redis
	redisReady=redisUtils.RedisInit(beego.AppConfig.String("serverPath"),
		beego.AppConfig.String("redisPort"),beego.AppConfig.String("auth"))//
	if !redisReady{
		beego.Error("redis connected error\n")
		os.Exit(1)
	}
	dbutils.InitMySQLDB(beego.AppConfig.String("mysqluser"),
	beego.AppConfig.String("mysqlpass"),
	beego.AppConfig.String("mysqlurls"),
	beego.AppConfig.String("mysqldb"))
	go server.TcpInit()
}

func main() {

	//初始化sql 配置文件
	utils3.Init("/static/cont/datasource.xml")
	s:=utils3.GetItemSql("USER_QUERY","GET_ALLUSER")
	b,_:=json.Marshal(s)
	fmt.Println(string(b))

	beego.Run()
}

