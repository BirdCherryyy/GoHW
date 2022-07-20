package main

import "fmt"

func main() {
	var a, b float32
	fmt.Println("Программа подсчета площади прямоугольника")
	fmt.Println("Введите значение стороны А")
	fmt.Scanln(&a)
	fmt.Println("Введите значение стороны B")
	fmt.Scanln(&b)
	fmt.Println("Площадь прямоугольника равна A * B = ", a, " * ", b, " = ", a*b)
}
