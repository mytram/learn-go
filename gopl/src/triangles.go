package main

import (
	"fmt"
	"strings"
)

type List struct {
	Elements []int
}

type Operator func(int) int

func (l *List) At(position int) int {
	return l.Elements[position]
}

func (l *List) Each(operator Operator) *List {
	newList := List{Elements: make([]int, len(l.Elements))}

	for i, v := range l.Elements {
		newList.Elements[i] = operator(v)
	}

	return &newList
}

func copySlice(src []int) []int {
	var dst []int

	for _, element := range src {
		dst = append(dst, element)
	}

	return dst
}

func main() {
	list := new(List)
	list.Elements = []int{1, 2, 3}

	newList := list.Each(func(a int) int { return a * 2 })

	fmt.Println(newList.Elements)

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
