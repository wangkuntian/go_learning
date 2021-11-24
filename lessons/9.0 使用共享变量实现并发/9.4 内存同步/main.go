package main

import (
	"fmt"
	"time"
)

func main() {
	var x, y int
	go func() {
		x = 1                   // A1
		fmt.Print("y:", y, " ") // A2
	}()
	go func() {
		y = 1                   // B1
		fmt.Print("x:", x, " ") // B2
	}()
	time.Sleep(1 * time.Second)
}
