package main

import (
	"exercices/area"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"
)

func main() {
	http.HandleFunc("POST /area", func(w http.ResponseWriter, r *http.Request) {
		_ = r.ParseForm()

		length := r.Form.Get("length")
		log.Printf("length: %v\n", length)
		width := r.Form.Get("width")
		log.Printf("width: %v\n", width)

		err := area.Main(strings.NewReader("meter\n"+length+"\n"+width+"\n"), w)
		if err != nil {
			log.Printf("err: %v\n", err)
			w.WriteHeader(http.StatusBadRequest)
		}
	})

	http.Handle("/", http.FileServer(http.Dir("area/")))

	fmt.Fprintln(os.Stdout, "listening...")
	err := http.ListenAndServe(":3333", nil)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
