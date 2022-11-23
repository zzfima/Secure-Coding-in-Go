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

	//log.Printf("%d bytes in %v", n, time.Since(start))
	fmt.Fprintf(w, "%d bytes digested in %v", n, time.Since(start))
}

// RunServerTimeoutsDDOS ...
func RunServerTimeoutsDDOS() {
	http.HandleFunc("/", handler)

	srv := http.Server{
		Addr:              ":8080",
		ReadTimeout:       1 * time.Second,
		WriteTimeout:      1 * time.Second,
		ReadHeaderTimeout: 2 * time.Second,
		IdleTimeout:       10 * time.Second,
	}

	if e := srv.ListenAndServe(); e != nil {
		log.Fatal(e)
	}
}
