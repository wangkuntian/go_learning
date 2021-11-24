package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"time"
)

func dirents2(dir string) []os.FileInfo {
	entries, err := ioutil.ReadDir(dir)
	if err != nil {
		fmt.Fprintf(os.Stderr, "du1: %v \n", err)
		return nil
	}
	return entries
}
func walkDir2(dir string, fileSizes chan<- int64) {
	for _, entry := range dirents2(dir) {
		if entry.IsDir() {
			subdir := filepath.Join(dir, entry.Name())
			walkDir2(subdir, fileSizes)
		} else {
			fileSizes <- entry.Size()
		}
	}

}

func printDiskUsage2(nfiles, nbytes int64) {
	fmt.Printf("%d files %.1f GB \n", nfiles, float64(nbytes)/1e9)
}

func main() {
	var verbose = flag.Bool("v", false, "show verbose progress message")
	flag.Parse()
	roots := flag.Args()
	if len(roots) == 0 {
		roots = []string{"."}
	}
	fileSizes := make(chan int64)
	go func() {
		for _, root := range roots {
			walkDir2(root, fileSizes)
		}
		close(fileSizes)
	}()
	var tick <-chan time.Time
	if *verbose {
		tick = time.Tick(500 * time.Millisecond)
	}

	var nfiles, nbytes int64
loop:
	for {
		select {
		case size, ok := <-fileSizes:
			if !ok {
				break loop
			}
			nfiles++
			nbytes += size

		case <-tick:
			printDiskUsage2(nfiles, nbytes)

		}
	}
	printDiskUsage2(nfiles, nbytes)
}
