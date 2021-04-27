package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {
	var w sync.WaitGroup
	w.Add(100)
	runtime.GOMAXPROCS(2)
	for i := 0; i < 100; i++ {
		go func() {
			fmt.Print(0)
			w.Done()
		}()
		fmt.Print(1)
	}
	w.Wait()
}
