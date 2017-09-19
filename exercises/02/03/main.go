package main

import "fmt"

func greatest(values ...int) int {
	max := values[0]
	for _, v := range values {
		if v > max {
			max = v
		}
	}
	return max
}

func main() {
	fmt.Println(greatest(11, 22, 44444, 22, 566, 346, 6545))
}
