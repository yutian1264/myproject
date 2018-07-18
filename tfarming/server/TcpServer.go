package server

import (
	"flag"
	"fmt"
	"time"
	"bufio"
	"net"
	"os"
	"encoding/json"
)

var hostser = flag.String("host", "", "host")
var portser = flag.String("port", "9999", "port")

type MsgSer struct {
	Data string `json:"data"`
	Type int    `json:"type"`
}

type Resp struct {
	Data string `json:"data"`
	Status int  `json:"status"`
}

func TcpInit()error{
	// 解析参数
	flag.Parse()
	var l net.Listener
	var err error
	// 监听
	l, err = net.Listen("tcp", *hostser+":"+*portser)
	if err != nil {
		defer func(err1 error)error{
			recover()
			fmt.Println("Error listening:", err)
			return err1
		}(err)
		panic(err)

	}
	defer l.Close()
	fmt.Println("Listening on " + *hostser + ":" + *portser)
	for {
		// 接收一个client
		conn, err := l.Accept()
		if err != nil {
			fmt.Println("Error accepting: ", err)
			os.Exit(1)
		}

		fmt.Printf("Received message %s -> %s \n", conn.RemoteAddr(), conn.LocalAddr())

		// 执行
		go handleRequest(conn)
	}
	return nil
}

// 处理接收到的connection
//
func handleRequest(conn net.Conn) {

	ipStr := conn.RemoteAddr().String()
	fmt.Println("client ip :" + ipStr)
	defer func() {
		fmt.Println("Disconnected :" + ipStr)
		conn.Close()
	}()

	// 构建reader和writer
	reader := bufio.NewReader(conn)
	writer := bufio.NewWriter(conn)

	for {
		// 读取一行数据, 以"\n"结尾
		//b, _, err := reader.ReadLine()
		b, err := reader.ReadString('\n')
		//b, err := reader.ReadByte()
		//fmt.Println("byte: ", b)
		if err != nil {
			fmt.Println("err: ", err)
			return
		}
		// 反序列化数据
		var msg MsgSer
		//json.Unmarshal(b, &msg)
		fmt.Println("GET ==>  data: ", string(b), " type: ", msg.Type)

		// 构建回复Msg
		resp := Resp{
			Data: time.Now().String(),
			Status: 200,
		}
		r, _ := json.Marshal(resp)

		writer.Write(r)
		writer.Write([]byte("\n"))
		writer.Flush()
	}

	fmt.Println("Done!")
}
