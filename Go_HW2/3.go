package main

import (
	"fmt"
	"os"
)

func main() {
	var a int64
	fmt.Println("Введите трехзначное число")
	fmt.Scanln(&a)
	if a >= 100 || a <= 999 {
		fmt.Println("Сотни = ", a/100)
		fmt.Println("Десятки = ", (a-(a/100)*100)/10)
		fmt.Println("Единицы = ", a-a/100*100-(a-(a/100)*100)/10*10)
	} else {
		fmt.Println("Некорректное число")
	}
	path, exists := os.LookupEnv("PATH")
	if exists {
		// Print the value of the environment variable
		fmt.Print(path)
	}
}
