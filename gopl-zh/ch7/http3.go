package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	db := database{"shoes": 500, "socks": 150}
	mux := http.NewServeMux()
	// http.HandlerFunc(db.list)是类型转换
	mux.Handle("/list", http.HandlerFunc(db.list))
	mux.Handle("/price", http.HandlerFunc(db.price))
	// 第二种方式
	// mux.HandleFunc("/list", db.list)
	// mux.HandleFunc("/price", db.price)
	log.Fatal(http.ListenAndServe("localhost:8000", mux))
}

type dollars float32

func (d dollars) String() string {
	return fmt.Sprintf("$%.2f", d)
}

type database map[string]dollars

func (db database) list(w http.ResponseWriter, r *http.Request) {
	for item, price := range db {
		fmt.Fprintf(w, "%s: %s\n", item, price)
	}
}

func (db database) price(w http.ResponseWriter, r *http.Request) {
	item := r.URL.Query().Get("item")
	price, ok := db[item]
	if !ok {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, "no such item: %q\n", item)
		return
	}
	fmt.Fprintf(w, "%s\n", price)
}
