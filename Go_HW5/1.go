package main

import (
	"fmt"
)
//1, 2, 3, 4, 5, 6...
//1, 2, 3, 5, 8, 13, 21, 34, 55, 89, 144
func main (){
	fmt.Println(Fib(2))
}

func Fib(Number int) int{
	if (Number < 2){
		return 1
	}else{
		return (Fib(Number - 1) + Fib(Number - 2))
	}
}