package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println("Программа подсчета диаметра и длины круга по заданному значению площади")
	var S, d, P float64
	fmt.Println("Введите значение площади круга")
	fmt.Scanln(&S)
	//S=Pi*R^2 -> R = sqrt(S/Pi) -> d = R*2
	d = float64(math.Sqrt(S / math.Pi))
	fmt.Println("Диаметр круга равен: ", d)
	P = 2 * math.Pi * (d / 2)
	fmt.Println("Длина круга равна: ", P)
}
