package main

import "fmt"

//带原型管理器的原型模式
func main() {
	cat1 := new(Cat)
	cat1.Name = "cat1"

	dog1 := new(Dog)
	dog1.Name = "dog1"
	//设置原型管理类
	manager := new(AnimnalProtoManager)
	manager.SetProto("cat", cat1)
	manager.SetProto("dog", dog1)

	//从原型管理类中获取所需的原型
	dog2, ok := manager.GetProto("dog").Clone().(IAnimal)
	if !ok {
		fmt.Println("转换失败")
	}
	dog2.Speak()

	cat2, ok := manager.GetProto("cat").Clone().(*Cat)
	if !ok {
		fmt.Println("转换失败")
	}
	fmt.Println(cat2.Name)
	cat2.Speak()

}

type IAnimal interface {
	Cloneable
	Speak()
}

func (dog *Dog) Speak() {
	fmt.Println("汪汪汪。。。。")
}

type Cat struct {
	Name   string
	Master *Person
}

func (cat *Cat) Clone() interface{} {
	newCat := new(Cat)
	newCat.Name = cat.Name
	newCat.Master = cat.Master //这里用的是浅拷贝
	return newCat
}

func (cat *Cat) Speak() {
	fmt.Println("喵喵喵^^^^^^")
}

//可以根据传入不同的参数来生成指定的狗原型，比如，田园犬
type AnimnalProtoManager struct {
	protoMap map[string]IAnimal
}

func (manager *AnimnalProtoManager) GetProto(class string) IAnimal {
	proto, ok := manager.protoMap[class]
	if ok {
		return proto
	}
	return nil
}

func (manager *AnimnalProtoManager) SetProto(class string, proto IAnimal) {
	if manager.protoMap == nil {
		manager.protoMap = make(map[string]IAnimal)
	}
	manager.protoMap[class] = proto
}
