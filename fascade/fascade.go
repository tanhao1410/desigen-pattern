package main

import (
	"fmt"
	"strings"
)

//外观模式:实现功能->公司前台相当于外观类，办事，里面有很多部门
func main() {
	cf := NewComponyFront()
	cf.ProcessProblem("hr的问题")
	cf.ProcessProblem("关于money的问题")
	//sync.WaitGroup{}.
	//sync.
}

//抽象外观模式
type IComponyFront interface {
	ProcessProblem(question string)
}

func NewComponyFront() *ComponyFront {
	cf := new(ComponyFront)
	cf.hr = new(HRDepartment)
	cf.money = new(MoneyDepartment)
	return cf
}

//公司前台--->外观类，给客户端提供功能
type ComponyFront struct {
	hr    *HRDepartment
	money *MoneyDepartment
}

//对外提供一个统一的办事方法
func (front *ComponyFront) ProcessProblem(question string) {
	if strings.Contains(question, "hr") {
		front.hr.Hire()
	} else if strings.Contains(question, "money") {
		front.money.GiveMoney()
	}
}

//hr部门
type HRDepartment struct {
}

//hr部门招人
func (department *HRDepartment) Hire() {
	fmt.Println("hr:我来看看还缺不缺人了")
}

//财务
type MoneyDepartment struct {
}

func (department *MoneyDepartment) GiveMoney() {
	fmt.Println("财务：我来看看还欠工资不？")
}
