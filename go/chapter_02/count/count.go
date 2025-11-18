package count

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net/http"
	"strconv"
)

func prompt(in io.Reader, out io.Writer) string {
	// fmt.Fprint(out, "What is the input string? ")
	s := bufio.NewScanner(in)

	var ans string
	for {
		s.Scan()
		if ans = s.Text(); ans != "" {
			break
		}
		// fmt.Fprintf(out, "Input is empty. Please enter a string. ")
	}

	return ans
}

func Count(in io.Reader, out io.Writer) {
	input := prompt(in, out)
	log.Printf("input: %v\n", input)
	// fmt.Fprintf(out, "%s has %d characters.\n", input, len(input))
	fmt.Fprint(out, len(input))
}

func Main(in io.Reader, out io.Writer) int {
	http.HandleFunc("POST /clicked", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("<div>content</div>"))
	})

	http.HandleFunc("POST /count", func(w http.ResponseWriter, r *http.Request) {
		_ = r.ParseForm()

		q := r.Form.Get("q")
		log.Printf("b: %v\n", q)
		io.WriteString(w, strconv.Itoa(len(q)))
		// Count(io.Reader(r.Body), w)
		// w.Write([]byte("<div>content</div>"))
	})

	http.Handle("/", http.FileServer(http.Dir("count/")))

	fmt.Fprintln(out, "listening...")
	// http.ListenAndServe(":3333", nil)
	return 0
}
