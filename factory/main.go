package main

import "fmt"

//工厂模式
func main() {
	//得到工厂
	factory := new(LenovoFactory)
	pc := factory.ProducePC()
	phone := factory.ProducePhone()
	pc.Compute()
	phone.Telecom()
}

//定义产品的接口（电脑）
type IPC interface {
	Compute()
}
type IPhone interface {
	Telecom()
}

//定义一个抽象电脑工厂
type IFactory interface {
	ProducePC() IPC
	ProducePhone() IPhone
}

//定义一个具体电脑工厂，工厂实现类
type LenovoFactory struct {
}

func (factory *LenovoFactory) ProducePC() IPC {
	return &LenovoPC{"Lenovo"}
}

func (factory *LenovoFactory) ProducePhone() IPhone {
	return &LenovoPhone{"Lenovo"}
}

//定义一个联想电脑，产品实现类 ，具体产品
type LenovoPC struct {
	Label string
}

func (pc *LenovoPC) Compute() {
	fmt.Println(pc.Label, "：运行。。。。。。")
}

//定义一个联想 手机实例
type LenovoPhone struct {
	Label string
}

func (phone *LenovoPhone) Telecom() {
	fmt.Println(phone.Label, "：打电话。。。。。。")
}
