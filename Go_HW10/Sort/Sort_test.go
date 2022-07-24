package sort

//Испотзование Example

//вывод отсортированной копии массива
func ExampleSortM() {
	var arr [10]int = [10]int{4, 8, 3, 2, 7, 1, 0, 6, 9, 5}
	err := sortM(arr)
	if err != nil {
		panic(err)
	}
}
