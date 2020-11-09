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

func main() {
	server1()
}

// server1: A simple echo server, returns the path
// component of the URL used to access the server
// e.g. http://localhost:8000/hello
// the response will be URL.Path = "/hello"
func server1() {
	http.HandleFunc("/", handler1) //each request calls handler
	http.HandleFunc("/count", counter) // returns the counts of requests made
	http.HandleFunc("/report", handler2) // reports on the request made
	http.HandleFunc("/gif", func(w http.ResponseWriter, r *http.Request) {

		//fmt.Println("GET params were:", r.URL.Query())

		// if only one expected
		cycles := r.URL.Query().Get("cycles")
		if cycles == "" {
			cycles = "5"
		}
		cyc, _ := strconv.Atoi(cycles)

		lissajous(w, cyc)
	})
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

// handler1 echoes the Path component of the requested URL
func handler1(w http.ResponseWriter, r *http.Request) {
	mu.Lock()
	count++
	fmt.Println("in handler1 in mu")
	mu.Unlock()
	fmt.Println("in handler1 after mu")
	fmt.Fprintf(w, "URL.Path = %s\n", r.URL.Path)
}

// counter echoes the number of counts so far
func counter(w http.ResponseWriter, r *http.Request) {
	mu.Lock()
	fmt.Fprintf(w, "Count %d\n", count)
	mu.Unlock()
}

// handler2: the handler function can report on the headers
// and form data that it receives, making the server useful
// for inspecting and debugging requests
func handler2(w http.ResponseWriter, r *http.Request) {

	// Prints the Method, URL, and Protocol
	fmt.Fprintf(w, "%s %s %s\n", r.Method, r.URL, r.Proto)

	// Goes through each value of the Header map
	// printing each key and value
	for k, v := range r.Header {
		fmt.Fprintf(w, "Header[%q] = %q\n", k, v)
	}

	// Prints the Host and Remote Address
	fmt.Fprintf(w, "Host = %q\n", r.Host)
	fmt.Fprintf(w, "RemoteAddr = %q\n", r.RemoteAddr)

	// ParseForm populates r.Form and r.PostForm.
	// For all requests, ParseForm parses the raw
	// query from the URL and updates r.Form.
	// writing it in the if statement shortens the
	// scope of err
	if err := r.ParseForm(); err != nil {
		log.Print(err)
	}

	//! DOES NOT POPULATE FORM IN A GET REQUEST, POST WORKS THO
	fmt.Println(r.Form)
	// Goes through each key and value of form
	for k, v := range r.Form {
		fmt.Fprintf(w, "Form[%q] = %q\n", k, v)
	}
}

// the same function from 1.4 but modified to take
// a different cycle value in params
func lissajous(out io.Writer, cyc int) {
	var palette = []color.Color{
		color.Black,
		color.RGBA{0, 255, 0, 255}, // Green
		color.RGBA{255, 0, 0, 255}, // Red
		color.RGBA{0, 0, 255, 255}, // Blue
	}

	cycles := cyc // number of complete x oscillator revolutions
	const (
		res     = 0.001 // angular resolution
		size    = 100   // image canvas covers [-size..+size]
		nframes = 128   // number of animation frames
		delay   = 5     // delay between frames in 10ms units
	)

	freq := rand.Float64() * 3.0 // relative frequency of y oscillator

	anim := gif.GIF{LoopCount: nframes} // gif struct, used to build animation

	phase := 0.0 // phase difference

	// iterate through each frame
	for i := 0; i < nframes; i++ {

		// create a rectangle, e.g. if size=100 then
		// dimensions are 201x201
		rect := image.Rect(0, 0, 2*size+1, 2*size+1)

		//Paletted image with the given width, height of the rect and palette
		img := image.NewPaletted(rect, palette)

		// randomly chooses a color from the palette for frame
		frameColor := uint8(rand.Intn(len(palette)-1) + 1)

		//Calculate each pixel of the oscillation in this current frame
		for t := 0.0; t < float64(cycles)*2*math.Pi; t += res {
			x := math.Sin(t)
			y := math.Sin(t*freq + phase)
			img.SetColorIndex(size+int(x*size+0.5), size+int(y*size+0.5),
				frameColor)
		}

		// adds 0.1 to phase
		phase += 0.1

		// append delay and image to the animation for this frame
		anim.Delay = append(anim.Delay, delay)
		anim.Image = append(anim.Image, img)
	}

	// encode gif to writer passing the reference to the struct
	gif.EncodeAll(out, &anim) // NOTE: ignoring encoding errors
}
