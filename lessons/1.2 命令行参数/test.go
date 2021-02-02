package main

import (
	"fmt"
	"os"
	"strings"
	"time"
)

func echo() {
	s, step := "", ""
	for _, arg := range os.Args[1:] {
		s += step + arg
		step = " "
	}
	fmt.Println(s)
}

func echo2() {
	fmt.Println(strings.Join(os.Args[1:], " "))
}

func echo3() {
	fmt.Println(strings.Join(os.Args, " "))
}

func echo4() {
	for index, arg := range os.Args {
		fmt.Println(index, arg)
	}
}

func main() {
	start := time.Now().Unix()
	echo()
	now := time.Now().Unix()
	fmt.Printf("cost: %d \n", (now - start))
	start = now
	echo2()
	now = time.Now().Unix()
	fmt.Printf("cost: %d \n", (now - start))
	start = now
	echo3()
	now = time.Now().Unix()
	fmt.Printf("cost: %d \n", (now - start))
	start = now
	echo4()
	now = time.Now().Unix()
	fmt.Printf("cost: %d \n", (now - start))

}
