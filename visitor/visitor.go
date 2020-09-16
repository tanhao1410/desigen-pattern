package main

import "fmt"

func main() {
	//初始化电脑城
	shop := new(ComputerShop)
	computers := []IComputer{new(GameComputer), new(ThinComputer)}
	shop.Computers = computers

	//两个顾客
	c1 := new(ManCustomer)
	c2 := new(WomanCustomer)

	//向两个顾客介绍产品
	shop.IndroductionComputer(c1)
	shop.IndroductionComputer(c2)
}

type ComputerShop struct {
	Computers []IComputer
}

func (shop *ComputerShop) IndroductionComputer(customer ICustomer) {
	for _, v := range shop.Computers {
		v.Accept(customer)
	}
}

type IComputer interface {
	Accept(customer ICustomer)
}

type GameComputer struct {
}

func (c *GameComputer) Accept(customer ICustomer) {
	fmt.Println(customer.EvaluateGameComputer(c))
}

type ThinComputer struct {
}

func (c *ThinComputer) Accept(customer ICustomer) {
	fmt.Println(customer.EvaluateThinComputer(c))
}

type ICustomer interface {
	//注意的是与java不同的是，golang中不支持重载，因此这里的方法名不同。
	EvaluateGameComputer(computer *GameComputer) string
	EvaluateThinComputer(computer *ThinComputer) string
}

type ManCustomer struct {
}

func (c *ManCustomer) EvaluateGameComputer(computer *GameComputer) string {
	return "男顾客评价游戏本：性能强，不错！"
}

func (c *ManCustomer) EvaluateThinComputer(computer *ThinComputer) string {
	return "男顾客评价轻薄本：性能差，打游戏不行！"
}

type WomanCustomer struct {
}

func (c *WomanCustomer) EvaluateGameComputer(computer *GameComputer) string {
	return "女顾客评价游戏本：又厚又丑！"
}

func (c *WomanCustomer) EvaluateThinComputer(computer *ThinComputer) string {
	return "女顾客评价轻薄本：这个好看！"
}
