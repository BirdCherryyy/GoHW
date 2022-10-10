package Fib

import "testing"

//реализованы T-тесты и B-тесты
func BenchmarkFirstFib(b *testing.B) {
	for i := 0; i < b.N; i++ {
		firstFib(3)
	}
}
func BenchmarkSecondFib(b *testing.B) {
	for i := 0; i < b.N; i++ {
		secondFib(3)
	}
}

func TestFirstFib(t *testing.T) {
	result := firstFib(3)
	if result != 3 {
		t.Errorf("asd:%d", result)
	}
}

func TestSecondFib(t *testing.T) {
	result := secondFib(3)
	if result != 3 {
		t.Errorf("asd:%d", result)
	}
}
