// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// Run with "web" command-line argument for web server.
// See page 13.
//!+main

// Lissajous generates GIF animations of random Lissajous figures.
package main

import (
	"image"
	"image/color"
	"image/gif"
	"io"
	"math"
	"math/rand"
	"os"
	"strconv"
)

//!-main
// Packages not needed by version in book.
import (
	"log"
	"net/http"
	"time"
)

//!+main
var white = color.RGBA{0x00, 0x00, 0x00, 0xff}
var black = color.RGBA{0xff, 0xff, 0xff, 0xff}
var green = color.RGBA{0x00, 0xff, 0x00, 0xff}
var red = color.RGBA{0xff, 0x00, 0x00, 0xff}
var blue = color.RGBA{0x00, 0x00, 0xff, 0xff}

var colors = [8]color.RGBA{white, black, green, red, blue}

func main() {
	//!-main
	// The sequence of images is deterministic unless we seed
	// the pseudo-random number generator using the current time.
	// Thanks to Randall McPherson for pointing out the omission.
	rand.Seed(time.Now().UTC().UnixNano())

	if len(os.Args) > 1 && os.Args[1] == "web" {
		http.HandleFunc("/", handler)
		log.Fatal(http.ListenAndServe("localhost:8080", nil))
		return
	}
	//!+main
	lissajous(os.Stdout, 5)
}

func handler(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		log.Print(err)
	}
	cycles := r.FormValue("cycles")
	if cycles != "" {
		c, err := strconv.Atoi(cycles)
		if err != nil {
			log.Print(err)
		}
		//fmt.Fprintf(w, "Form["cycles"] = %v\n", cycles)
		lissajous(w, float64(c))
	} else {
		lissajous(w, 5.0)
	}
}

func lissajous(out io.Writer, cycles float64) {
	const (
		res     = 0.001 // angular resolution
		size    = 100   // image canvas covers [-size..+size]
		nframes = 64    // number of animation frames
		delay   = 8     // delay between frames in 10ms units
	)

	lineIndex := rand.Intn(5)
	canvasIndex := 5 - lineIndex
	palette := []color.Color{colors[canvasIndex], colors[lineIndex]}

	freq := rand.Float64() * 3.0 // relative frequency of y oscillator
	anim := gif.GIF{LoopCount: nframes}
	phase := 0.0 // phase difference

	for i := 0; i < nframes; i++ {
		rect := image.Rect(0, 0, 2*size+1, 2*size+1)
		img := image.NewPaletted(rect, palette)
		for t := 0.0; t < cycles*2*math.Pi; t += res {
			x := math.Sin(t)
			y := math.Sin(t*freq + phase)
			img.SetColorIndex(size+int(x*size+0.5), size+int(y*size+0.5),
				uint8(lineIndex))
		}
		phase += 0.1
		anim.Delay = append(anim.Delay, delay)
		anim.Image = append(anim.Image, img)
	}
	gif.EncodeAll(out, &anim) // NOTE: ignoring encoding errors
}

//!-main
