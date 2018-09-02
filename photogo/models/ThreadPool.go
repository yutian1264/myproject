package models

import "fmt"

type Pool struct {
	Queue chan func() error;
	RuntineNumber int;
	Total int;

	Result chan error;
	FinishCallback func();
}

//初始化
func (self *Pool) Init(runtineNumber int,total int)  {
	self.RuntineNumber = runtineNumber;
	self.Total = total;
	self.Queue = make(chan func() error, total);
	self.Result = make(chan error, total);
}

func (self *Pool) Start()  {
	//开启 number 个goruntine
	for i:=0;i<self.RuntineNumber;i++ {
		go func() {
			for {
				task,ok := <-self.Queue
				if !ok {
					break;
				}
				err := task();
				self.Result <- err;
			}
		}();
	}

	//获取每个任务的处理结果
	for j:=0;j<self.Total;j++ {
		res,ok := <-self.Result;
		if !ok {
			break;
		}
		if res != nil {
			fmt.Println(res);
		}
	}

	//结束回调函数
	if self.FinishCallback != nil {
		self.FinishCallback();
	}
}

//关闭
func (self *Pool) Stop()  {
	close(self.Queue);
	close(self.Result);
}

func (self *Pool) AddTask(task func() error)  {
	self.Queue <- task;
}

func (self *Pool) SetFinishCallback(fun func())  {
	self.FinishCallback = fun;
}

