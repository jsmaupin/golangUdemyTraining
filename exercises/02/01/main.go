package main

import "fmt"

func half(val int) (int, bool) {
	return val / 2, val%2 == 0
}

func main() {
	for i := 0; i < 100; i++ {
		var halved, isEven = half(i)
		fmt.Printf("input=%d half=%d isEven=%v\n", i, halved, isEven)
	}
}
