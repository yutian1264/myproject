package models

import (
	"io/ioutil"
	"os"
	"strings"
	"log"
	"github.com/rwcarlsen/goexif/exif"
)

type Images struct {
	Name string
	City string
	CreateTime string
	Path string
	Lat float64
	Lng float64

}

func OpenFilesAndGetMsg(path string, suffix []string) (images []Images, err error) {
	var files = make([]string, 0, 10)
	dir, err := ioutil.ReadDir(path)
	if err != nil {
		return nil, err
	}
	files = recursively(path, suffix, files, dir)

	images=GetExifMess(files)
	return images, nil
}

//递归获取图片信息
func recursively(path string, suffix, files []string, dir []os.FileInfo) []string {
	PthSep := string(os.PathSeparator)
	for _, fi := range dir {
		if fi.IsDir() { // 忽略目录
			d2, _ := ioutil.ReadDir(path + PthSep + fi.Name())
			files = recursively(path + PthSep + fi.Name(), suffix, files, d2)

		}
		name := strings.ToUpper(fi.Name())
		for _, s := range suffix {
			if strings.HasSuffix(name, s) { //匹配文件
				files = append(files, path+PthSep+fi.Name())
			}
		}

	}
	return files
}

func GetExifMess(imgs []string)[]Images {
	var imgList =make([]Images,0,10)
	if len(imgs) > 0 {
		for _, s := range imgs {
			var images=Images{}
			//fmt.Println(s)
			images.Path=s
			f1, err := os.Open(s)
			images.Name=f1.Name()
			if err != nil {
				recover()
				f1.Close()
				log.Println("error:",s)
				imgList=append(imgList,images )
				continue
			}
			defer f1.Close()
			exif, err := exif.Decode(f1)
			if err != nil {
				log.Println("Decode:",s)
				recover()
				f1.Close()
				continue
			}
			//t:=exif.Get("DateTime")
			lat,lng,err:=exif.LatLong()
			if err!=nil{
				log.Println("LatLong:",err)
			}
			images.Lat=lat
			images.Lng=lng
			//todo 根据经纬度获取城市信息
			a, e := exif.Get("DateTime")
			if e != nil {
				recover()
				f1.Close()
				log.Println("DateTime:",s)
				imgList=append(imgList,images )
				continue
			}
			images.CreateTime=a.String()
			imgList=append(imgList,images )
		}

	}
	return imgList
}

func main() {
	suffix := []string{".PNG", ".JPG"}
	OpenFilesAndGetMsg("D:\\image", suffix)
}
