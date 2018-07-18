package server

import (
	"net"
	"sync"
)
var mux sync.RWMutex
var tmList [] net.Conn=make([]net.Conn,1000000)
type Terminal struct {

}
func(t *Terminal)DeleteItem(c net.Conn)bool{


	return false
}

func(t *Terminal)Add(c net.Conn){
	mux.RLock()
	tmList=append(tmList, c)
	mux.RUnlock()

}


