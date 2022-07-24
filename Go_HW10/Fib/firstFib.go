package Fib

func firstFib(Number int) int {
	if Number < 2 {
		return 1
	} else {
		return (firstFib(Number-1) + firstFib(Number-2))
	}
}

func secondFib(Number int) int {
	var mp map[int]int = map[int]int{}
	for k := range mp {
		if k == Number {
			return mp[k]
		}
	}
	if Number < 2 {
		mp[Number] = 1
		return 1
	} else {
		mp[Number] = (secondFib(Number-1) + secondFib(Number-2))
		return (secondFib(Number-1) + secondFib(Number-2))
	}
}
