package main

import (
	"exercices/simplemath"
	"fmt"
	"log"
	"net/http"
	"strings"
)

func main() {
	http.HandleFunc("POST /count", func(w http.ResponseWriter, r *http.Request) {
		err := r.ParseForm()
		if err != nil {
			fmt.Fprintf(w, "error parsing form: %s", err)
			w.WriteHeader(http.StatusBadRequest)
		}

		first := r.Form.Get("first")
		second := r.Form.Get("second")
		rr := strings.NewReader(first + "\n" + second + "\n")

		err = simplemath.Main(rr, w)
		if err != nil {
			fmt.Fprintf(w, "error calculating result: %s", err)
			w.WriteHeader(http.StatusBadRequest)
		}
	})

	http.Handle("/", http.FileServer(http.Dir("simplemath/")))

	log.Println("listening...")
	err := http.ListenAndServe(":3333", nil)
	log.Printf("err: %v\n", err)
}
