package main

import (
	"fmt"
)

var mp map[int]int = map[int]int{}

//1, 2, 3, 4, 5, 6...
//1, 2, 3, 5, 8, 13, 21, 34, 55, 89, 144
func main() {
	fmt.Println(Fib(5))
	fmt.Println(Fib(5))
	for key, value := range mp {
		fmt.Println("Key:", key, "Value:", value)
	}
}

func Fib(Number int) int {
	for k := range mp {
		if k == Number {
			return mp[k]
		}
	}
	if Number < 2 {
		mp[Number] = 1
		return 1
	} else {
		mp[Number] = (Fib(Number-1) + Fib(Number-2))
		return (Fib(Number-1) + Fib(Number-2))
	}
}
