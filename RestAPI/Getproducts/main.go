package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/go-chi/chi"
)

// func getProducts(w http.ResponseWriter, r *http.Request){
// 	fmt.Println("Hello")
// }
func main() {
	r := chi.NewRouter()
	r.Get("/products", func(w http.ResponseWriter, r *http.Request) {
		resp, err := http.Get("http://localhost:3000/products")
		if err != nil {
			panic(err)
		}
		defer resp.Body.Close()

		responseData, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			log.Fatal(err)
		}
		responseString := string(responseData)
		fmt.Println(responseString)
	})

	http.ListenAndServe(":3333", r)
}
