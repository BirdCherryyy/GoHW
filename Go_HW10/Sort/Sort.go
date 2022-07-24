package sort

import "fmt"

//вывод отсортированной копии массива
func sortM(arr [10]int) error {
	mem := 0
	for i := 0; i < len(arr)-1; i++ { //цикл который запоминает текущее значения по i
		for i2 := i; i2 >= 0; i2-- { // цикл, который работает с массивом при помощи i2
			if arr[i2+1] < arr[i2] { //условия свапа элементов в массиве
				mem = arr[i2+1]
				arr[i2+1] = arr[i2]
				arr[i2] = mem
			}
		}
		fmt.Println(arr)
	}
	return nil
}
