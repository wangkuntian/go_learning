package main

import "fmt"

func main() {
	var i int = 0
	var a *int
	a = &i
	fmt.Println("Hello World", i, a)
}
