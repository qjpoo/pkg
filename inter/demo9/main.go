package main

import "fmt"

type Usb interface {
	start()
	end()
}

type Mouse struct {
	name string
}

type FlashDisk struct {
	name string
}

func (m Mouse)start()  {
	fmt.Println(m.name, " start ...")
}

func (m Mouse)end()  {
	fmt.Println(m.name,  " end ...")
}

func (f FlashDisk)start()  {
	fmt.Println(f.name, " start ...")
}

func (f FlashDisk)end()  {
	fmt.Println(f.name, " end ...")
}

func (f FlashDisk)delete()  {
	fmt.Println(f.name, "delete data ...")
}


// 测试接口
func TestUsb(usb Usb)  {
	usb.start()
	usb.end()
}

func main() {
	// interface
	var i Usb
	m := Mouse{"Dell鼠标"}
	i = m
	i.start()
	i.end()

	TestUsb(m)

	f := FlashDisk{"金士顿的U盘"}
	TestUsb(f)
	i = f
	i.start()
	i.end()
	f.delete()  // f有删除功能，而usb没有
	//i.delete() // usb没有delete方法
	// i.name  // usb不能访问name字段，因为name字段只有mouse和flashdisk有

}
