package models

import (
	"time"
)

//获取当前时间
func GetCurrentTime() string {
	return time.Now().Format("2006-01-02 15:04:05")
}

//时间对比
func TimeChange(t1, t2 string) bool {
	format := "2006-01-02 15:04:05"
	a, _ := time.Parse(format, t1)
	b, _ := time.Parse(format, t2)
	return a.After(b.Add(300 * time.Second))
}
