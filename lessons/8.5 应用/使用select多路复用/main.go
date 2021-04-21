package main

import (
	"fmt"
	"os"
	"time"
)

func launch() {
	fmt.Println("Launched")
}

//func main() {
//	abort := make(chan struct{})
//	go func() {
//		os.Stdin.Read(make([]byte, 1))
//		abort <- struct{}{}
//	}()
//	fmt.Println("Commencing countdown. Press return to abort.")
//	select {
//	case <-time.After(10 * time.Second):
//	case <-abort:
//		fmt.Println("Launch aborted")
//		return
//	}
//	launch()
//}

//func main() {
//	ch := make(chan int, 1)
//	for i := 0; i < 10; i++ {
//		select {
//		case x := <-ch:
//			fmt.Println(x)
//		case ch <- i:
//		}
//	}
//}

func main() {
	abort := make(chan struct{})
	go func() {
		os.Stdin.Read(make([]byte, 1))
		abort <- struct{}{}
	}()
	fmt.Println("Commencing countdown. Press return to abort.")
	tick := time.Tick(1 * time.Second)
	for countdown := 10; countdown > 0; countdown-- {
		fmt.Println(countdown)
		select {
		case <-tick:
		case <-abort:
			fmt.Println("Launch aborted")
			return
		}
	}
	launch()
}
