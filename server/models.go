package main

import (
	"net"
	"sync"
)

type Msg struct {
	Message string
	From    string
	To      string
}

type ConnMngr struct {
	ConnMap map[string]net.Conn
	mutex   sync.RWMutex
}

func NewConnMngr() *ConnMngr {
	cm := make(map[string]net.Conn)
	return &ConnMngr{
		ConnMap: cm,
		mutex:   sync.RWMutex{},
	}
}


func (cm *ConnMngr) Add(clientAddress string, conn net.Conn) {
	cm.mutex.Lock()
	defer cm.mutex.Unlock()
	cm.ConnMap[clientAddress] = conn
}

func (cm *ConnMngr) Remove(clientAddress string) {
	cm.mutex.Lock()
	defer cm.mutex.Unlock()
	delete(cm.ConnMap, clientAddress)
}

func (cm *ConnMngr) Get(clientAddress string) net.Conn {
	cm.mutex.RLock()
	defer cm.mutex.RUnlock()
	return cm.ConnMap[clientAddress]
}

func (cm *ConnMngr) All() []net.Conn {
	cm.mutex.RLock()
	defer cm.mutex.RUnlock()
	ret := []net.Conn{}
	for _, c := range cm.ConnMap {
		ret = append(ret, c)
	}
	return ret
}

func (cm *ConnMngr) Close() []net.Conn {
	cm.mutex.Lock()
	defer cm.mutex.Unlock()
	ret := make([]net.Conn, len(cm.ConnMap))
	for _, c := range cm.ConnMap {
		_ = c.Close()
	}
	return ret
}