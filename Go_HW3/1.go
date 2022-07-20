package main

import (
	"fmt"
	"math"
)

func main() {
	var A, B float64
	var str string
	fmt.Println("Введите числа A и B")
	fmt.Scanln(&A, &B)
	fmt.Println("Что с ними сделать? (+, -, /, *, **, !(для факториала будет использоваться только А))")
	fmt.Scanln(&str)
	switch str {
	case "+":
		fmt.Printf("%.2f + %.2f = %.2f", A, B, A+B)
	case "-":
		fmt.Printf("%.2f - %.2f = %.2f", A, B, A-B)
	case "*":
		fmt.Printf("%.2f * %.2f = %.2f", A, B, A*B)
	case "/":
		if B == 0 {
			fmt.Println("Ошибка деления на 0")
		} else {
			fmt.Printf("%.2f / %.2f = %.2f", A, B, A/B)
		}
	case "**":
		fmt.Printf("%.2f ** %.2f = %.2f", A, B, math.Pow(A, B))
	case "!":
		fmt.Printf("%.2f", factorial(A))
	}
}

func factorial(x float64) (result float64) {
	if x == 0 {
		result = 1
	} else {
		result = x * factorial(x-1)
	}
	return
}
