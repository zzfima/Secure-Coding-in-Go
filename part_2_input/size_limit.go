package part2

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

func logNoLimiterHandler(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	if a, e := io.ReadAll(r.Body); e != nil {
		log.Fatal(e)
	} else {
		fmt.Fprintf(w, "bytes stored in DB %d", len(a))
	}
}

// RunServerSizeNoLimiting ...
func RunServerSizeNoLimiting() {
	http.HandleFunc("/", logNoLimiterHandler)
	http.ListenAndServe(":8080", nil)
}

func logYesLimiterHandler(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	if a, e := io.ReadAll(io.LimitReader(r.Body, 100)); e != nil {
		log.Fatal(e)
	} else {
		fmt.Fprintf(w, "bytes stored in DB %d", len(a))
	}
}

// RunServerSizeYesLimiting ...
func RunServerSizeYesLimiting() {
	http.HandleFunc("/", logYesLimiterHandler)
	http.ListenAndServe(":8080", nil)
}
