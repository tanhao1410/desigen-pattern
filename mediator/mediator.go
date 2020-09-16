package main

import "fmt"

func main() {
	//中介者
	mediator := NewConcreteMediator()

	//具体同事
	collegue1 := &ConcreteCollegue{
		Name:     "同事1",
		Mediator: mediator,
	}
	collegue2 := &ConcreteCollegue{
		Name:     "同事2",
		Mediator: mediator,
	}
	collegue3 := &ConcreteCollegue{
		Name:     "同事3",
		Mediator: mediator,
	}
	collegue4 := &ConcreteCollegue{
		Name:     "同事4",
		Mediator: mediator,
	}
	//注册到中介者上
	mediator.Register(collegue1)
	mediator.Register(collegue2)
	mediator.Register(collegue3)
	mediator.Register(collegue4)

	//同事1向其它同事发消息
	mediator.Relay(collegue1, "同事1想要传达的信息")
}

//抽象同事类
type ICollegue interface {
	SetMediator(mediator IMediator)
	Receive(message string)
	Send(message string)
}

//抽象中介类
type IMediator interface {
	Register(collegue ICollegue)
	Relay(collegue ICollegue, message string)
}

//具体同事类
type ConcreteCollegue struct {
	Name     string
	Mediator IMediator
}

func (collegue *ConcreteCollegue) Receive(message string) {
	fmt.Println("我是", collegue.Name, "收到了消息：", message)
}

func (collegue *ConcreteCollegue) Send(message string) {
	collegue.Mediator.Relay(collegue, message)
}

func (collegue *ConcreteCollegue) SetMediator(mediator IMediator) {
	collegue.Mediator = mediator
	mediator.Register(collegue)
}

//具体中介类
type ConcreteMediator struct {
	Collegues map[ICollegue]Empty //此处仅仅是把map当做set来用
}

//一个空的结构体
type Empty struct {
}

func (mediator *ConcreteMediator) Register(collegue ICollegue) {
	mediator.Collegues[collegue] = Empty{}
}

func (mediator *ConcreteMediator) Relay(collegue ICollegue, message string) {
	for k, _ := range mediator.Collegues {
		if k != collegue {
			k.Receive(message)
		}
	}
}

//工厂方法模式
func NewConcreteMediator() IMediator {
	mediator := new(ConcreteMediator)
	mediator.Collegues = make(map[ICollegue]Empty)
	return mediator
}
