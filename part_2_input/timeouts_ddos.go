package part2

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"time"
)

func handler(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	start := time.Now()
	n, e := io.Copy(io.Discard, r.Body)

	if e != nil {
		http.Error(w, "can not copy", http.StatusBadRequest)
		return
	}

	log.Printf("%d bytes in %v", n, time.Since(start))
	fmt.Fprintf(w, "%d bytes digested", n)
}

// RunServerTimeoutsDDOS ...
func RunServerTimeoutsDDOS() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}
