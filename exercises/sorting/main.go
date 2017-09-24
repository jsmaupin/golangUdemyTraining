package main

import (
	"fmt"
	"sort"
)

type people []string

func (p people) Len() int {
	return len(p)
}

func (p people) Less(i, j int) bool {
	return p[i] < p[j]
}

func (p people) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}

func main() {

	// (1)
	studyGroup := people{"Zeno", "John", "Al", "Jenny"}
	sort.Sort(studyGroup)
	fmt.Println(studyGroup)
	sort.Sort(sort.Reverse(studyGroup))
	fmt.Println(studyGroup)

	// (2)
	s := []string{"Zeno", "John", "Al", "Jenny"}
	ss := sort.StringSlice(s)
	ss.Sort()
	fmt.Println(s)
	sort.Sort(sort.Reverse(ss))
	fmt.Println(s)

	// (3)
	n := []int{7, 4, 8, 2, 9, 19, 12, 32, 3}
	ns := sort.IntSlice(n)
	ns.Sort()
	fmt.Println(n)
	sort.Sort(sort.Reverse(ns))
	fmt.Println(ns)
}
