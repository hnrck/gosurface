package main

import (
	"github.com/hnrck/gosurface/pkg/surface"
	"log"
	"net/http"
	"math"
)

func main() {
	http.HandleFunc("/", surfaceHandler)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

func surfaceHandler(w http.ResponseWriter, _ *http.Request) {
	const (
		width, height = 600, 320
		cells         = 100
		xyrange       = 30.0
		angle         = math.Pi / 6
	)
	s := surface.Surface{
		Width:   width,
		Height:  height,
		Cells:   cells,
		XYRange: xyrange,
		Angle:   angle,
	}
	w.Header().Set("Content-Type", "image/svg+xml")
	s.GeneratePolygons()
	s.Writer(w)
}
