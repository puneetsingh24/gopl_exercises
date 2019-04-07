package main

import "fmt"

func main() {
	a := [...]int{1, 2, 3, 4, 5, 6, 7}
	reverse(&a)
	fmt.Println(a) // [7 6 5 4 3 2 1 0]
}

func reverse(s *[7]int) {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}
