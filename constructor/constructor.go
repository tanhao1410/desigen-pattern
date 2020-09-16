package main

import "fmt"

//建造者模式，电脑由
//包括：产品：电脑PC由主机mainframes、显示器screen、鼠标Mouse组成
//抽象建造者：电脑厂
//具体建造者：具体电脑厂家
//指挥者：导购
func main() {
	//两个厂家
	lenove := new(Lenove)
	asus := new(Asus)

	//销售
	saleMan := &SaleMan{asus}
	saleMan2 := &SaleMan{lenove}

	fmt.Printf("info：%#v\n", saleMan.SalePC())
	fmt.Printf("info:%#v", saleMan2.SalePC())
}

type SaleMan struct {
	Factory IPCFactory //哪个商家的
}

func (saleMan *SaleMan) SalePC() *PC {
	return saleMan.Factory.ProduceComputer()
}

type PC struct {
	Mainframes string
	Screen     string
	Mouse      string
}

//着重的是在于生产不同的组件，然后组装在一起。这是和工厂模式的区分。如果不加入这些生产部分组件的 方法，和工厂模式无异
//在java中有抽象类，而golang中没有，因此，生产总
type IPCFactory interface {
	ProduceComputer() *PC
	ProduceMainframes() string
	ProduceScreen() string
	ProduceMouse() string
}

//联想
type Lenove struct {
}

func (factory *Lenove) ProduceComputer() *PC {
	pc := new(PC)
	pc.Mainframes = factory.ProduceMainframes()
	pc.Mouse = factory.ProduceMouse()
	pc.Screen = factory.ProduceScreen()
	return pc
}
func (factory *Lenove) ProduceMainframes() string {
	return "联想的主机"
}
func (factory *Lenove) ProduceScreen() string {
	return "联想的屏幕"
}
func (factory *Lenove) ProduceMouse() string {
	return "联想的鼠标"
}

//华硕
type Asus struct {
}

func (factory *Asus) ProduceComputer() *PC {
	pc := new(PC)
	pc.Mainframes = factory.ProduceMainframes()
	pc.Mouse = factory.ProduceMouse()
	pc.Screen = factory.ProduceScreen()
	return pc
}
func (factory *Asus) ProduceMainframes() string {
	return "华硕的主机"
}
func (factory *Asus) ProduceScreen() string {
	return "华硕的屏幕"
}
func (factory *Asus) ProduceMouse() string {
	return "华硕的鼠标"
}
