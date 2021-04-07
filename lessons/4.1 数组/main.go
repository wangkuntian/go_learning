package main

import "fmt"

func reverse(s []int) {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}

func main() {
	l := [...]int{0, 1, 2, 3, 4, 5}
	c := l[2:3]
	fmt.Println(c)
	fmt.Println(cap(c), len(c))
	s := []int{0, 1, 2, 3, 4, 5}
	fmt.Printf("%T\n", s)
	reverse(s[:2])
	fmt.Println(s)
	reverse(s[2:])
	fmt.Println(s)
	reverse(s)
	fmt.Println(s)
}
