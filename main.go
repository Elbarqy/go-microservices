package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		log.Println("hello world")
		d, err := ioutil.ReadAll(r.Body)
		if err != nil {
			http.Error(w, "oops", http.StatusBadRequest)
			// w.WriteHeader(http.StatusBadRequest)
			// w.Write([]byte("Oooopse"))
			return
		}
		// log.Printf("Data %s", d)
		fmt.Println(w, "Hello %s", d)
	})

	http.ListenAndServe("localhost:9090", nil)
}
