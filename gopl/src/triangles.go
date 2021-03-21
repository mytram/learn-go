package main

import (
	"fmt"
	"strings"
)

func copySlice(src []int) []int {
	var dst []int

	for _, element := range src {
		dst = append(dst, element)
	}

	return dst
}

func main() {
	// there are 20 rods
	var rods [20]int

	for i := 1; i <= 20; i++ {
		rods[i-1] = i
	}

	var slice1 []int
	slice2 := copySlice(rods[2:])

	newSlice := append(slice1, slice2...)

	printSlice(slice1)
	printSlice(slice2)
	printSlice(newSlice)

	printSlice(slice1)
	printSlice(newSlice)

	printSlice(rods[0:])
}

func printSlice(s []int) {
	var ss []string

	for _, i := range s {
		ss = append(ss, fmt.Sprintf("%d", i))
	}

	fmt.Println(strings.Join(ss, ","))
}
