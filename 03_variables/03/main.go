package main

import "fmt"

var a = [4]int{1, 2, 3}

func main() {
	b := [3]int{}

	fmt.Println(a)
	fmt.Printf("%T\n", a)

	fmt.Println(b)
	fmt.Printf("%T\n", b)
}
