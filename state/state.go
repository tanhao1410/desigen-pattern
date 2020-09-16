package main

import "fmt"

//状态设计模式：
//1.抽象状态类（饥饿状态）<------具体状态类（饱了，没饱）
//2.环境类(人）
func main() {
	p := new(Person)
	startState := new(Hungry)
	p.HungryState = startState

	p.Eat()
	p.Work()
	p.Work()
	p.Eat()
	p.Eat()

}

//人，环境类
type Person struct {
	HungryState IHungry
}

func (person *Person) Eat() {
	if person.HungryState.IsHungry() {
		fmt.Println("eatting.........")
		//改变状态
		person.HungryState = new(NoHungry)
	} else {
		fmt.Println("already baole!!!")
	}
}

func (person *Person) Work() {
	if person.HungryState.IsHungry() {
		fmt.Println("i'm hungry ,no work!")
	} else {
		fmt.Println("do work...........")
		//改变状态
		person.HungryState = new(Hungry)
	}
}

//抽象状态类 饥饿接口
type IHungry interface {
	IsHungry() bool
}

//具体状态类
type Hungry struct {
}

func (hungry *Hungry) IsHungry() bool {
	return true
}

type NoHungry struct {
}

func (noHungry *NoHungry) IsHungry() bool {
	return false
}
