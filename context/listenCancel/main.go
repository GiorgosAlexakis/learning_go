package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"
)

func main() {

	log.Fatal(http.ListenAndServe("localhost:8000", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		fmt.Fprintf(os.Stdout, "Started Processing Request\n")
		Duration := 2 * time.Second
		select {
		case <-time.After(Duration):
			w.Write([]byte("process requested"))
		case <-ctx.Done():
			fmt.Fprintf(os.Stderr, "request cancelled")
		}
	})))
}
