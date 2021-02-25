package main

import (
	"fmt"
)

type chargerHead interface {
	UseRoundPlug()
}
type RussiaChargerHead struct {

}

func (r*RussiaChargerHead)	UseRoundPlug() {
	fmt.Println("Connect with round plug")
}

type ChineseChargerHead struct {

}

func (c*ChineseChargerHead) UseStrongPlug() {
	fmt.Println("Connect with strong plug")
}

type americanChargerHead struct {

}

func (a* americanChargerHead) UseStrongPlug () {
	fmt.Println("Connect with american strong plug")

}

type strongAdapter struct {
	changer *ChineseChargerHead
}

func (s*strongAdapter) UseRoundPlug()  {
	s.changer.UseStrongPlug()
}
type aStrongAdapter struct {
	americanCharger *americanChargerHead
}

func (a*aStrongAdapter)UseRoundPlug () {
	a.americanCharger.UseStrongPlug()
}
type person struct {

}

func (p*person) UseRoundPlugToMyDevice(char chargerHead) {
	char.UseRoundPlug()
}
func main() {
	person:=&person{}
	roundPlug:=&RussiaChargerHead{}
	person.UseRoundPlugToMyDevice(roundPlug)
	changer:=&ChineseChargerHead{}
	changerMachine:=&strongAdapter{
		changer,
	}
	person.UseRoundPlugToMyDevice(changerMachine)
	americanCharger:=&americanChargerHead{}
	americanChargerMachine:=&aStrongAdapter{
		americanCharger: americanCharger,
	}
	person.UseRoundPlugToMyDevice(americanChargerMachine)


}