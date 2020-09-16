package main

import (
	"fmt"
)

func main() {

	//组装责任链
	logHandler := new(LogHandler)
	authoHandler := new(AuthoHandler)
	authoHandler.SetNext(logHandler)

	//责任链间传递的信息
	info := RequestInfo{Name: "test"}

	//客户调用责任链的入口责任
	authoHandler.Handle(info)

}

type RequestInfo struct {
	Name string
}

//责任链接口，所有的责任链实现者都应实现该接口
type Handler interface {
	//责任链中的处理方法
	Handle(info RequestInfo)
	//设置下一个
	SetNext(handler Handler)
	//获取下一个
	GetNext() (next Handler, ok bool)
}

//记录日志
type LogHandler struct {
	next Handler
}

func (handler *LogHandler) Handle(info RequestInfo) {
	fmt.Println("日志记录：", info.Name)
	//接着调用后续处理
	if next, ok := handler.GetNext(); ok {
		next.Handle(info)
	}
}

func (handler *LogHandler) SetNext(next Handler) {
	handler.next = next
}

func (handler *LogHandler) GetNext() (next Handler, ok bool) {
	if handler.next != nil {
		return handler.next, true
	}
	return nil, false
}

//权限验证
type AuthoHandler struct {
	next Handler
}

func (handler *AuthoHandler) Handle(info RequestInfo) {
	fmt.Println("权限验证：", info.Name)
	if next, ok := handler.GetNext(); ok {
		next.Handle(info)
	}
}

func (handler *AuthoHandler) SetNext(next Handler) {
	handler.next = next
}

func (handler *AuthoHandler) GetNext() (next Handler, ok bool) {
	if handler.next != nil {
		return handler.next, true
	}
	return nil, false
}
