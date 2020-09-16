package main

import "fmt"

//命令模式，客户端通过调用者，传递不同的命令，然后不同的接受者对此进行处理
func main() {
	invoker := NewInvoker(NewShutDownTV())
	invoker.Call()
}

//调用者
type Invoker struct {
	command ICommand
}

func NewInvoker(command ICommand) *Invoker {
	invoker := new(Invoker)
	invoker.command = command
	return invoker
}

func (invoker *Invoker) Call() {
	if invoker.command != nil {
		invoker.command.Execute()
	}
}

//命令接口
type ICommand interface {
	Execute()
}

type ShutDownTV struct {
	TV *TV
}

func NewShutDownTV() *ShutDownTV {
	commandTV := new(ShutDownTV)
	commandTV.TV = NewTV()
	return commandTV
}

func (command *ShutDownTV) Execute() {
	command.TV.ShutDown()
}

type TV struct {
}

func NewTV() *TV {
	return new(TV)
}

func (tv *TV) ShutDown() error {
	fmt.Println("关闭电视")
	return nil
}
