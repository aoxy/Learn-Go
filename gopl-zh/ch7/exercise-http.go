package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"sync"
)

func main() {
	db := database{table: map[string]dollars{"shoes": 500, "socks": 150}}
	http.HandleFunc("/list", db.list)
	http.HandleFunc("/price", db.price)
	http.HandleFunc("/update", db.update)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

type dollars float32

func (d dollars) String() string {
	return fmt.Sprintf("$%.2f", d)
}

type database struct {
	table map[string]dollars
	mutex sync.Mutex
}

func (db database) list(w http.ResponseWriter, r *http.Request) {
	for item, price := range db.table {
		fmt.Fprintf(w, "%s: %s\n", item, price)
	}
}

func (db database) price(w http.ResponseWriter, r *http.Request) {
	item := r.URL.Query().Get("item")
	price, ok := db.table[item]
	if !ok {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, "no such item: %q\n", item)
		return
	}
	fmt.Fprintf(w, "%s\n", price)
}

func (db database) update(w http.ResponseWriter, r *http.Request) {
	item := r.URL.Query().Get("item")
	price, err := strconv.ParseFloat(r.URL.Query().Get("price"), 32)
	if err != nil || price < 0 {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "invalid price: %f\n", price)
		return
	}
	db.mutex.Lock()
	db.table[item] = dollars(price)
	db.mutex.Unlock()
}
