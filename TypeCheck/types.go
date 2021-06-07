package main

import "fmt"

func Do(i interface{}) {
	switch v := i.(type) {
	case int:
		fmt.Println("Int ", v)
	case string:
		fmt.Println("String")

	}
}

func main() {
	i := 10
	st := "Siva"

	Do(i)
	Do(st)
}
