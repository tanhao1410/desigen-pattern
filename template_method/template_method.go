package main

import "fmt"

func main() {
	factory := new(ComputerFactory)
	factory.RealFactory = new(LenovoFactory)
	fmt.Println(factory.ProduceComputer())
}

type IComputerFactory interface {
	ProduceComputer() string
	BuyAllUnits()
	Assemble() string
}

//相当于抽象类
type ComputerFactory struct {
	RealFactory IComputerFactory //包含一个真正的实现的引用
}

//在这里写“抽象类”的共同的方法，实际上调用的是具体的实现来完成的。在java中，是通过this来引用具体子类，而在此处，是通过RealFactory属性
func (factory *ComputerFactory) ProduceComputer() string {
	factory.RealFactory.BuyAllUnits()
	return factory.RealFactory.Assemble()
}

//相当于具体实现子类，匿名组合父类，然后重写在java中属于抽象的方法
type LenovoFactory struct {
	ComputerFactory
}

func (factory *LenovoFactory) BuyAllUnits() {
	fmt.Println("Lenovo 去购买了组件")
}
func (factory *LenovoFactory) Assemble() string {
	return "Lenovo computer"
}
