package main

import "fmt"

func main() {
	result := (true && false) || (false && true) || !(false && false)
	fmt.Println("The result is", result)
}
