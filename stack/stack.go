package main

import (
	"fmt"
)

func Push(stack *[]string, s string) {
	*stack = append(*stack, s)
}

func Pop() {
	fmt.Println(stack[len(stack)-1])
	stack = stack[:len(stack)-1]
}

var stack []string

func main() {

	Push(&stack, "Hello")
	Push(&stack, "World")
	Push(&stack, "..!")

	fmt.Println(stack[0:])
	Pop()
	fmt.Println(stack[0:])

}
