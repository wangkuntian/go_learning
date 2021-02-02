package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
)

func fetch() {
	for _, url := range os.Args[1:] {
		res, err := http.Get(url)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Fetch: %v \n", err)
			os.Exit(1)
		}
		b, err := ioutil.ReadAll(res.Body)
		res.Body.Close()
		if err != nil {
			fmt.Fprintf(os.Stderr, "Fetch: reading %s: %v \n", url, err)
			os.Exit(1)
		}
		fmt.Printf("%s", b)
	}
}

func fetch2() {
	for _, url := range os.Args[1:] {
		if strings.HasPrefix(url, "http://") || strings.HasPrefix(url, "https://") {
			res, err := http.Get(url)
			if err != nil {
				fmt.Fprintf(os.Stderr, "Fetch: %v \n", err)
				os.Exit(1)
			}
			_, err = io.Copy(os.Stdout, res.Body)
			res.Body.Close()
			if err != nil {
				fmt.Fprintf(os.Stderr, "Fetch: reading %s: %v \n", url, err)
			}
			fmt.Printf("\n %s \n", res.Status)
		} else {
			continue
		}
	}
}

func main() {
	fetch2()
}

// cat 1.txt | xargs go run main.go
