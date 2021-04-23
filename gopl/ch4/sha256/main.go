package main

import (
	"crypto/sha256"
	"fmt"
)

func countDiffBits(p1 []byte, p2 []byte) int {
	size1 := len(p1)
	size2 := len(p2)

	if size1 != size2 {
		return -1
	}

	var s1 = make([]string, size1)
	var s2 = make([]string, size2)

	for i, v := range p1 {
		s1[i] = fmt.Sprintf("%08b", v)
	}

	for i, v := range p2 {
		s2[i] = fmt.Sprintf("%08b", v)
	}

	count := 0
	for i := 0; i < size1; i++ {
		bits1 := []byte(s1[i])
		bits2 := []byte(s2[i])

		for j := 0; j < 8; j++ {
			if bits1[j] != bits2[j] {
				count++
			}
		}
	}

	return count
}

func updateSlice(s []int, pos int, value int) {
	s[pos] = value
}

func main() {
	c1 := sha256.Sum256([]byte("X"))
	c2 := sha256.Sum256([]byte("X"))

	fmt.Println(countDiffBits(c1[0:], c2[0:]))

	fmt.Printf("%08b\n", c1)
	fmt.Printf("%08b\n", c2)

	a := [...]int{1, 2, 3, 4}

	updateSlice(a[1:], 0, 10)

	fmt.Printf("%d\n", a)
}
