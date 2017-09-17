package main

import "fmt"

func main() {
	var smallNumber int
	var largeNumer int

	fmt.Print("Enter a small number: ")
	fmt.Scan(&smallNumber)

	fmt.Print("Enter a large number: ")
	fmt.Scan(&largeNumer)

	fmt.Printf("%d %% %d = %d\n", largeNumer, smallNumber, largeNumer%smallNumber)
}
