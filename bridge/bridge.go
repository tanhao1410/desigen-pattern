package main

import "fmt"

func main() {

	windows := new(WindowsLaptop)
	windows2 := new(WindowsLaptop)
	linux := new(LinuxLaptop)

	game := new(GameType)
	thin := new(ThinType)

	//可以根据不同的维度进行自由组合了
	windows.LaptopType = game
	windows2.LaptopType = thin
	linux.LaptopType = game

	windows.Run()
	fmt.Println("================================")
	windows2.Run()
	fmt.Println("================================")
	linux.Run()
}

type LaptopType interface {
	GetType() string
}

type ILaptop interface {
	Run()
}

type GameType struct {
}

func (laptopType *GameType) GetType() string {
	return "游戏本"
}

type ThinType struct {
}

func (laptopType *ThinType) GetType() string {
	return "轻薄本"
}

type WindowsLaptop struct {
	LaptopType LaptopType
}

func (laptop *WindowsLaptop) Run() {
	if laptop.LaptopType == nil {
		fmt.Println("该笔记本未设置具体类型")
	} else {
		fmt.Println("笔记本类型为：", laptop.LaptopType.GetType())
	}
	fmt.Println("windows正在运行。。。。。。")

}

type LinuxLaptop struct {
	LaptopType LaptopType
}

func (laptop *LinuxLaptop) Run() {
	if laptop.LaptopType == nil {
		fmt.Println("该笔记本未设置具体类型")
	} else {
		fmt.Println("笔记本类型为：", laptop.LaptopType.GetType())
	}
	fmt.Println("linux正在运行。。。。。。")

}
