package controllers

import (
	"github.com/astaxie/beego"
	"fmt"
	"encoding/json"
	"sky/com/utils/dbUtils"
	"path"
	"log"
	"os"
	"io"
	"sky/com/utils/file"
	"mime/multipart"
	"sky/com/utils/upload"
)

type MainController struct {
	beego.Controller
}

func (c *MainController) URLMapping() {
	c.Mapping("conndb", c.InitDb)
	c.Mapping("upload", c.UploadHandle)
	c.Mapping("upload2", c.UploadHandle2)
}

func (c *MainController) Get() {
	c.TplName="tweenlit_demo.html"
	//c.TplName="upload.html"
}
// @router /addphoto [post]
func (this *MainController) Upload() {

	f, h, _ := this.GetFile("myfile")	//获取上传的文件

	f.Close()
	//	this.SaveToFile("image", path+ h.Filename)	//存文件    WaterMark(path)	//给文件加水印
	this.SaveToFile("image", path.Join("static/uploadfile",h.Filename))
	this.TplName ="index.html"

	// 根据字段名获取表单文件
}
//@router /upload [post]
func (c *MainController)UploadHandle() {
	c.TplName ="index.html"
	// 根据字段名获取表单文件

	formFile, header, err := c.GetFile("myfile")
	singFileUpload(formFile, header, err,"")
	if err != nil {
		log.Printf("Get form file failed: %s\n", err)
		return
	}
}

//@router /upload2 [post]
func (c *MainController)UploadHandle2() {
	c.TplName ="index.html"
	fmt.Println("0000000000000000000")
	// 根据字段名获取表单文件
//	models.GetPayLoad(c)
	go upload.UploadBreakPoint(c.Ctx.Request,"file","static/upload/")
}






func singFileUpload(formFile multipart.File, header *multipart.FileHeader, err error,savePath string) (bool,error){


	defer formFile.Close()
	b,err:=file.IsPathExists(savePath)
	if !b{
		err=os.MkdirAll(savePath,os.ModePerm)
		if err!=nil{
			log.Printf("Create directory error: %s\n", err)
			return false,err
		}
	}

	destFile, err := os.Create(savePath + header.Filename)
	if err != nil {
		log.Printf("Create failed: %s\n", err)
		return false,err
	}
	defer destFile.Close()
	_, err = io.Copy(destFile, formFile)
	if err != nil {
		log.Printf("Write file failed: %s\n", err)
		return false,err
	}
	return true,nil
}

func(this *MainController)Post(){
	ss:=this.GetString("param")
	var dat map[string]interface{}
	err := json.Unmarshal([]byte(ss), &dat)
	if  err == nil {
		fmt.Println(dat["a"])
	} else {
		fmt.Println(err)
	}
	fmt.Println("测试返回信息")

	this.Ctx.WriteString("{\"id\":\"abcd\"}")
	//this.Data["json"]="{\"ObjectId\":\"abcd\"}"
	//this.ServeJSONP()

}
// @router /conndb [post]
func (this *MainController)InitDb(){
	fmt.Println("conn db")

	m:=utils.Query("select* from users")

	d,err:=json.Marshal(m)
	if err!=nil{
		fmt.Println(" map to json error")
	}

	this.Ctx.WriteString(string(d))
}

