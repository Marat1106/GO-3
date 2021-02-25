package main

import "fmt"

type computer interface {
	boot()
	shutDown()
}
type DesktopComputer struct {
	power
	sound
	mainBoard
	cpu
}
type power struct {
}

func (p power)powerOn()  {
	fmt.Println("Power On")
}
func (p power) powerOff()  {
	fmt.Println("Power Off")
}
type sound struct {
}

func (s sound) soundOn()  {
	fmt.Println("Sound on")

}
func (s sound)soundOff()  {
	fmt.Println("sound off")
}

type mainBoard struct {}

func (m mainBoard) powerOn()  {
	fmt.Println("MainBoard power on")
}
func (m mainBoard)powerOff()  {
	fmt.Println("MainBoard power off")
}
type cpu struct {}

func (c cpu) powerOn() {
	fmt.Println("CPU power on")
}
func (c cpu) powerOff()  {
	fmt.Println("CPU power off")

}
func (dc DesktopComputer) boot()  {
	fmt.Println("Computer is opening")
	dc.power.powerOn()
	dc.cpu.powerOn()
	dc.mainBoard.powerOn()
	dc.sound.soundOn()
	fmt.Println("Computer was started successfully")
}
func (dc DesktopComputer) shutDown()  {
	fmt.Println("The computer is shutting down")
	dc.sound.soundOff()
	dc.mainBoard.powerOff()
	dc.cpu.powerOff()
	dc.power.powerOff()
}

type Laptop struct {
	power
	sound
	mainBoard
	cpu
}

func (l Laptop) Boot() {
	fmt.Println("Laptop is opening")
	l.power.powerOn()
	l.mainBoard.powerOn()
	l.cpu.powerOn()
	l.sound.soundOn()
	fmt.Println("opening successfully")
}

func (l Laptop) Shutdown() {
	fmt.Println("laptop starting close")
	l.sound.soundOff()
	l.cpu.powerOff()
	l.mainBoard.powerOff()
	l.power.powerOff()

	fmt.Println("shut down successfully")
}
type client struct {
}

func (c client)computer(t string) *Facade {
	return &Facade{}
}

type Facade struct {
	Desktop *DesktopComputer
	Laptop *Laptop

}


func main()  {
	laptop:=&Laptop{}
	desktop:=&DesktopComputer{}
	com:=(&client{}).computer("l")
	//if client use laptop
	com.Laptop=laptop
	com.Laptop.Boot()
	com.Laptop.Shutdown()
	//if client use desktop
	ad:=(&client{}).computer("dc")
	ad.Desktop=desktop
	ad.Desktop.boot()
	ad.Desktop.shutDown()
}