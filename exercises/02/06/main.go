package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

func main() {
	for x := 1; true; x++ {
		x2Set := strings.Split(strconv.Itoa(2*x), "")
		sort.Strings(x2Set)
		golden := strings.Join(x2Set, "")

		success := true
		for m := 3; m <= 6; m++ {
			set := strings.Split(strconv.Itoa(m*x), "")
			sort.Strings(set)
			test := strings.Join(set, "")
			if test != golden {
				success = false
				break
			}
		}

		if success {
			fmt.Println("The anser is:", x)
			fmt.Println()
			fmt.Println("Proof:")
			for i := 2; i <= 6; i++ {
				fmt.Printf("%d * %d = %d\n", i, x, i*x)
			}
			break
		}
	}
}

/*
Project Euler Problem 52

It can be seen that the number, 125874, and its double, 251748, contain exactly the same digits, but in a
different order.
Find the smallest positive integer, x, such that 2x, 3x, 4x, 5x, and 6x, contain the same digits.

The solution here is to use 2x as the golden value. If 3x, 4x, etc. all contain the same digits, then will contain
the same digits that are in 2x.
To see if they contain the same digits as 2x, the method is to:
1. Turn the number into a string
2. Split the string into characters
3. Sort the characters
4. Rejoin into a string
5. Perform a string compare
*/
