package main

import "fmt"

var num int = 1

func Even(ch chan bool) {
	if num%2 == 0 {
		fmt.Println(num)
		num++
		ch <- true
	}
}

func Odd(ch chan bool) {
	if num%2 != 0 {
		fmt.Println(num)
		num++
		ch <- true
	}
}

func main() {
	ch := make(chan bool)
	for i := 0; i < 10; i++ {
		go Odd(ch)
		<-ch

		go Even(ch)

		<-ch
	}
}
