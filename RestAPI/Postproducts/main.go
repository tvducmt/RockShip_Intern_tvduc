package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/jmcvetta/napping"
)

func main() {
	r := chi.NewRouter()
	r.Get("/post_product", func(w http.ResponseWriter, r *http.Request) {
		s := napping.Session{}
		h := &http.Header{}
		h.Set("X-Custom-Header", "myvalue")
		s.Header = h

		var jsonStr = []byte(`
			{
	 			"name": "Ipone 9",
	 			"price": 500,
		 		"status": true
			}`)

		var data map[string]json.RawMessage
		err := json.Unmarshal(jsonStr, &data)
		if err != nil {
			fmt.Println(err)
		}
		url := "http://localhost:3000/products"
		resp, err := s.Post(url, &data, nil, nil)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("response Status:", resp.Status())
		fmt.Println("response Headers:", resp.HttpResponse().Header)
		fmt.Println("response Body:", resp.RawText())
	})

	http.ListenAndServe(":3333", r)
}
