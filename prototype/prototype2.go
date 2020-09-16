package main

import (
	"fmt"
)

func main2() {
	person1 := new(Person)
	person1.Age = 12
	person1.Nation = "China"
	person1.Name = "test"
	dog1 := new(Dog)
	dog1.Name = "dog1"
	dog1.Master = person1

	dog2, ok := dog1.Clone().(*Dog) //注意这里是指针，因为方法返回的是指针类型
	if !ok {
		fmt.Println("转换失败。。。。")
	}
	fmt.Println(dog2.Name)
	fmt.Println(dog2.Master.Name)
}

type Dog struct {
	Name   string
	Master *Person
}

type Cloneable interface {
	Clone() interface{}
}

func (dog *Dog) Clone() interface{} {
	newDog := new(Dog)
	newDog.Name = dog.Name
	newDog.Master = dog.Master //这里用的是浅拷贝
	return newDog
}
