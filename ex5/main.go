package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Это главная страница!\n"))
}

func ProductsHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("это страница товаров!\n"))
}

func ArticleHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("это статья!\n"))
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", HomeHandler)
	r.HandleFunc("/products", ProductsHandler)
	r.HandleFunc("/articles/{category}/{id:[0-9]+}", ArticleHandler)

	http.Handle("/", r)
	err := http.ListenAndServe(":8080", nil)
	fmt.Println(err)
}
