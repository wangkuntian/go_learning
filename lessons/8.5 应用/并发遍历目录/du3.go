package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"sync"
)

var sema = make(chan struct{}, 20)

func dirents3(dir string) []os.FileInfo {
	sema <- struct{}{}
	defer func() { <-sema }()
	entries, err := ioutil.ReadDir(dir)
	if err != nil {
		fmt.Fprintf(os.Stderr, "du1: %v \n", err)
		return nil
	}
	return entries
}
func walkDir3(dir string, n *sync.WaitGroup, fileSizes chan<- int64) {
	defer n.Done()
	for _, entry := range dirents3(dir) {
		if entry.IsDir() {
			n.Add(1)
			subdir := filepath.Join(dir, entry.Name())
			go walkDir3(subdir, n, fileSizes)
		} else {
			fileSizes <- entry.Size()
		}
	}

}

func printDiskUsage3(nfiles, nbytes int64) {
	fmt.Printf("%d files %.1f GB \n", nfiles, float64(nbytes)/1e9)
}

func main() {
	flag.Parse()
	roots := flag.Args()
	if len(roots) == 0 {
		roots = []string{"."}
	}
	fileSizes := make(chan int64)
	var n sync.WaitGroup
	for _, root := range roots {
		n.Add(1)
		go walkDir3(root, &n, fileSizes)
	}
	go func() {
		n.Wait()
		close(fileSizes)
	}()
	var nfiles, nbytes int64
	for size := range fileSizes {
		nfiles++
		nbytes += size
	}
	printDiskUsage3(nfiles, nbytes)
}
