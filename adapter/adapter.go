package main

import "fmt"

func main() {
	computer := new(Computer)

	udisk := new(UDisk)
	hardDisk := new(HardDisk)
	box := new(HardDiskBox)

	box.HardDisk = hardDisk

	computer.Usb1 = udisk
	computer.Usb2 = box

	fmt.Println(computer.ReadUSB1())
	fmt.Println(computer.ReadUSB2())

}

type USB interface {
	ReadUSBDevice() string
}

type Computer struct {
	Usb1 USB
	Usb2 USB
}

func (c *Computer) ReadUSB1() string {
	return c.Usb1.ReadUSBDevice()
}

func (c *Computer) ReadUSB2() string {
	return c.Usb2.ReadUSBDevice()
}

type UDisk struct {
}

func (u *UDisk) ReadUSBDevice() string {
	return "这里是U盘，我是里面的数据。"
}

//硬盘，没有USB接口，无法差在电脑usb口上用
type HardDisk struct {
}

func (h *HardDisk) ReadHardDisk() string {
	return "这里是硬盘，我是里面的数据"
}

//硬盘盒，适配器，含有一个硬盘的引用，实现 USB接口，可以插在了电脑上面
type HardDiskBox struct {
	HardDisk *HardDisk
}

func (b *HardDiskBox) ReadUSBDevice() string {
	return b.HardDisk.ReadHardDisk()
}
