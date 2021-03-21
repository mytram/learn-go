package main

import (
	"fmt"
)

func main() {
	a := [...]int{1, 2, 3}

	fmt.Printf("the length of a is %d\n", len(a))

	slice := a[1:]

	fmt.Printf("the length slice is %d\n", len(slice))
}
