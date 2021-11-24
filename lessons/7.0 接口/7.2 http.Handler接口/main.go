package main

import (
	"fmt"
	"log"
	"net/http"
)

type dollars float32

func (d dollars) String() string {
	return fmt.Sprintf("$%.2f", d)
}

type database map[string]dollars

//func (db database) ServeHTTP(w http.ResponseWriter, req *http.Request) {
//
//	switch req.URL.Path {
//	case "/list":
//		for item, price := range db {
//			fmt.Fprintf(w, "%s: %s \n", item, price)
//		}
//	case "/price":
//		item := req.URL.Query().Get("item")
//		price, ok := db[item]
//		if !ok {
//			w.WriteHeader(http.StatusNotFound)
//			fmt.Fprintf(w, "no such item: %q \n", item)
//			return
//		}
//		fmt.Fprintf(w, "%s \n", price)
//	default:
//		w.WriteHeader(http.StatusNotFound)
//		fmt.Fprintf(w, "no such page: %s \n", req.URL)
//	}
//}
//
//func main() {
//	db := database{"shoes": 50, "socks": 5}
//	log.Fatal(http.ListenAndServe("localhost:8000", db))
//
//}

func (db database) list(writer http.ResponseWriter, request *http.Request) {
	for item, price := range db {
		fmt.Fprintf(writer, "%s: %s \n", item, price)
	}
}

func (db database) price(writer http.ResponseWriter, request *http.Request) {
	item := request.URL.Query().Get("item")
	price, ok := db[item]
	if !ok {
		writer.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(writer, "no such item: %q \n", item)
		return
	}
	fmt.Fprintf(writer, "%s \n", price)
}

//func main() {
//	db := database{"shoes": 50, "socks": 5}
//	mux := http.NewServeMux()
//	mux.Handle("/list", http.HandlerFunc(db.list))
//	mux.Handle("/price", http.HandlerFunc(db.price))
//	log.Fatal(http.ListenAndServe("localhost:8000", mux))
//}

func main() {
	db := database{"shoes": 50, "socks": 5}
	http.HandleFunc("/list", db.list)
	http.HandleFunc("/price", db.price)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}
