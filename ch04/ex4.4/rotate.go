package main

import "fmt"

func main() {
	a := []int{1, 2, 3, 4, 5, 6}
	a = rotateL(a, 2)
	fmt.Println(a)
}

func rotateL(s []int, num int) []int {
	if num > 0 && num <= len(s) {
		temp := s[:num]
		s = s[num:]
		s = append(s, temp...)
	}
	return s
}
