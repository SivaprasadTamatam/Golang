package main

import "fmt"

type Animal struct {
}

func (a *Animal) Eat() {
	fmt.Println("Animal Eat")
}

func (a *Animal) Run() {
	fmt.Println("Animal Run")
}

func (a *Animal) Sleep() {
	fmt.Println("Animal Eat")
}

type legs struct {
}

func (l *legs) NoofLegs() {
	fmt.Println("4 Legs")
}

type Dog struct {
	item int
	Animal
	legs
}

func (d *Dog) Eat() {
	fmt.Println("Dog Eat")
}

func main() {
	var d Dog
	d.item = 10
	fmt.Println(d.item)
	d.Eat()
	d.Run()
	d.NoofLegs()
}
