package main

import "testing"

func BenchmarkDayOnePartOne(b *testing.B) {
	for n := 0; n < b.N; n++ {
		dayOnePartOne()
	}
}

func BenchmarkDayOnePartTwo(b *testing.B) {
	for n := 0; n < b.N; n++ {
		dayOnePartTwo()
	}
}
