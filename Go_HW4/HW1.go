package main

import (
	"fmt"
)

func main() {
	fmt.Println("Сортировка вставками")
	var arr [10]int = [10]int{4, 8, 3, 2, 7, 1, 0, 6, 9, 5} //инициация массива
	fmt.Println("Изначальный массив", arr)
	var mem int
	for i := 0; i < len(arr)-1; i++ { //цикл который запоминает текущее значения по i
		for i2 := i; i2 >= 0; i2-- { // цикл, который работает с массивом при помощи i2
			if arr[i2+1] < arr[i2] { //условия свапа элементов в массиве
				mem = arr[i2+1]
				arr[i2+1] = arr[i2]
				arr[i2] = mem
			}
		}
		fmt.Println(arr) //вывод i-го этапа преобразования массива
	}
}
