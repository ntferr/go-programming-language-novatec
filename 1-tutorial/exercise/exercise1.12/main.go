package main

import (
	"fmt"
	"image"
	"image/color"
	"image/gif"
	"io"
	"log"
	"math"
	"math/rand"
	"net/http"
	"strconv"
	"sync"
)

var mu sync.Mutex
var count int
var pallet = []color.Color{
	color.RGBA{0, 0, 0, 1},       // Black
	color.RGBA{255, 255, 255, 1}, // White
	color.RGBA{128, 0, 128, 1},   // Purple
	color.RGBA{144, 238, 144, 1}, // Green
}

func main() {
	http.HandleFunc("/", handler)
	http.HandleFunc("/count", counter)
	http.HandleFunc("/lissajous", lissajousHandler)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "%s %s %s\n", r.Method, r.URL, r.Proto)
	for k, v := range r.Header {
		fmt.Fprintf(w, "Header[%q] = %q\n", k, v)
	}
	fmt.Fprintf(w, "Host = %q\n", r.Host)
	fmt.Fprintf(w, "RemoteAddr = %q\n", r.RemoteAddr)
	if err := r.ParseForm(); err != nil {
		log.Print(err)
	}
	for k, v := range r.Form {
		fmt.Fprintf(w, "Form[%q] = %q\n", k, v)
	}
}

func counter(w http.ResponseWriter, r *http.Request) {
	mu.Lock()
	count++
	fmt.Fprintf(w, "Count %d\n", count)
	mu.Unlock()
}

func lissajousHandler(w http.ResponseWriter, r *http.Request) {
	var paramValue int = 5
	var err error = nil

	params, ok := r.URL.Query()["cycle"]
	if ok {
		paramValue, err = strconv.Atoi(params[0])
		if err != nil {
			log.Printf("failed to convert string param to int: [%v]", err)
		}
	}

	lissajous(w, paramValue)
}

func lissajous(out io.Writer, paramValue int) {
	const (
		res     = 0.001
		size    = 100
		nframes = 64
		delay   = 8
	)
	cycles := float64(paramValue)
	freq := rand.Float64() * 3.0
	anim := gif.GIF{LoopCount: nframes}
	phase := 0.0

	for i := 0; i < nframes; i++ {
		rect := image.Rect(0, 0, 2*size+1, 2*size+1)
		img := image.NewPaletted(rect, pallet)

		for t := 0.0; t < cycles*2*math.Pi; t += res {
			x := math.Sin(t)
			y := math.Sin(t*freq + phase)
			randColor := rand.Intn(len(pallet)-1) + 1

			img.SetColorIndex(size+int(x*size+0.5), size+int(y*size+0.5), uint8(randColor))
		}
		phase += 0.1
		anim.Delay = append(anim.Delay, delay)
		anim.Image = append(anim.Image, img)
	}
	gif.EncodeAll(out, &anim)
}
