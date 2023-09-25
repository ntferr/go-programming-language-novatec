package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strconv"
)

var db = database{
	"shoes": 50,
	"socks": 5,
}

func main() {
	http.HandleFunc("/list", db.list)
	http.HandleFunc("/create", db.create)
	http.HandleFunc("/update", db.update)
	http.HandleFunc("/price", db.price)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

type dollars float32

func (d dollars) String() string { return fmt.Sprintf("$%.2f", d) }

type database map[string]dollars

func (db database) create(w http.ResponseWriter, req *http.Request) {
	item := req.URL.Query().Get("item")
	if item == "" {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "digit a item for add into database")
		return
	}

	priceString := req.URL.Query().Get("price")
	if priceString == "" {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "digit a price for add into database")
		return
	}
	price, err := strconv.ParseFloat(priceString, 32)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "digit a valid price")
		return
	}

	db[item] = dollars(price)

	fmt.Fprintf(w, "added %s: %s", item, db[item])

}

func (db database) list(w http.ResponseWriter, req *http.Request) {
	templ := template.Must(template.New("").Parse(`
{{range $key, $value := .}}
Product: {{$key}} Value: {{$value}}
{{end}}
	`))

	err := templ.Execute(w, db)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "an error ocurred to get list of products: %s", err.Error())
		return
	}
}

func (db database) price(w http.ResponseWriter, req *http.Request) {
	item, price, ok := searchIntoQueryRequest(req)
	if !ok {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, "no such item: %q\n", item)
		return
	}
	fmt.Fprintf(w, "%s\n", price)
}

func (db database) update(w http.ResponseWriter, req *http.Request) {
	item, _, ok := searchIntoQueryRequest(req)
	if !ok {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, "no such item: %q\n", item)
		return
	}

	priceString := req.URL.Query().Get("price")
	if priceString == "" {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "digit a price for add into database")
		return
	}
	priceUpdated, err := strconv.ParseFloat(priceString, 32)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "digit a valid price")
		return
	}

	db[item] = dollars(priceUpdated)

	fmt.Fprintf(w, "%s: %s\n", item, db[item])
}

func searchIntoQueryRequest(req *http.Request) (string, dollars, bool) {
	item := req.URL.Query().Get("item")
	price, ok := db[item]
	return item, price, ok
}
