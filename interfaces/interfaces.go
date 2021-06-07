package main

import "fmt"

type device interface {
	getType() string
	getCapacity() float64
}

type penDrive struct {
}

func (p *penDrive) getType() string {
	return "Pen Drive"
}

func (p *penDrive) getCapacity() float64 {
	return 12.0 * 1024 * 1024
}

func main() {
	var de device
	de = &penDrive{}

	d := penDrive{}
	fmt.Println(de.getType())

	fmt.Println(d.getCapacity())

}
