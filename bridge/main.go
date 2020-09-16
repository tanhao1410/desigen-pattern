package main

import "fmt"

//桥接模式
//通过组合的方式（而不是继承），扩展了原有的接口，而不用改动已经实现了原有接口的大量实现类，在新的扩展实现中，还可以利用原来的实现
func main() {
	var origin OriginFunc = new(OriginFuncImpl1)

	var extend ExtenderFunc = new(ExtendedImpl)
	extend.SetOriginImpl(origin)

	extend.ExtenderFunc1("这是扩展方法")
	fmt.Println(extend.GetOriginFunc().OriginFunc1()) //这是原有接口的功能
}

//原始功能接口
type OriginFunc interface {
	OriginFunc1() (string, bool)
}

//原始的一个实现类，系统中可能存在多个类似的
type OriginFuncImpl1 struct {
}

func (this *OriginFuncImpl1) OriginFunc1() (string, bool) {
	return "原始", true
}

//扩展接口
type ExtenderFunc interface {
	//Super OriginFunc
	ExtenderFunc1(s string) bool
	SetOriginImpl(originFunc OriginFunc)
	GetOriginFunc() OriginFunc
}

//扩展的实现类
type ExtendedImpl struct {
	super OriginFunc // 通过引用此来达到继承原来接口的能力
}

func (this *ExtendedImpl) SetOriginImpl(originFunc OriginFunc) {
	this.super = originFunc
}

func (this *ExtendedImpl) ExtenderFunc1(s string) bool {
	fmt.Println("扩展方法执行了")
	return false
}

func (this *ExtendedImpl) GetOriginFunc() OriginFunc {
	return this.super
}
