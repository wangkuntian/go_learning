package main

import (
	"fmt"
	"log"
	"net/http"
	"sync"
)

func server() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "URL.Path = %q \n", r.URL.Path)
	})
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func server2() {
	var mu sync.Mutex
	var count int
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		mu.Lock()
		count++
		mu.Unlock()
		fmt.Fprintf(w, "URL.Path = %q \n", r.URL.Path)
	})
	http.HandleFunc("/count", func(w http.ResponseWriter, r *http.Request) {
		mu.Lock()
		fmt.Fprintf(w, "Count: %d \n", count)
		mu.Unlock()
	})
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func server3() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "%s %s %s \n", r.Method, r.URL, r.Proto)
		for k, v := range r.Header {
			fmt.Fprintf(w, "Header[%q] = %q \n", k, v)
		}
		fmt.Fprintf(w, "Host = %q \n", r.Host)
		fmt.Fprintf(w, "RemoteAddr = %q \n", r.RemoteAddr)
		if err := r.ParseForm(); err != nil {
			log.Println(err)
		}
		for k, v := range r.Form {
			fmt.Fprintf(w, "Form[%q] = %q \n", k, v)
		}
	})
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func main() {
	//server()
	//server2()
	server3()
}
