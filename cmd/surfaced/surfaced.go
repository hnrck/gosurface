package main

import (
	"fmt"
	"github.com/hnrck/gosurface/pkg/surface"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", surfaceHandler)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

func surfaceHandler(w http.ResponseWriter, _ *http.Request) {
	surface := surface.Surface{}
	w.Header().Set("Content-Type", "image/svg+xml")
	fmt.Fprintf(w, "%s", surface)
}
