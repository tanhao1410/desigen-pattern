package main

import "fmt"

func main() {
	//生成一个普通的车
	car := new(Car)
	car.Run()

	carriage := new(ConcreteCarriage)
	carriage.car = car

	carriage.Run()
	carriage.CarryCargo()
}

type ICar interface {
	Run()
}
type Car struct {
}

func (car *Car) Run() {
	fmt.Println("开车。。。")
}

type ICarCarriage interface {
	ICar
	CarryCargo()
}

type ConcreteCarriage struct {
	car ICar
}

func (c *ConcreteCarriage) CarryCargo() {
	fmt.Println("带车厢的车 载货。。。")
}

func (c *ConcreteCarriage) Run() {
	//加入自己的逻辑
	fmt.Print("带车厢车 ")
	//调用原来的方法
	c.car.Run()
}
